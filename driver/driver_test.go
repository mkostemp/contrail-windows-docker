package driver

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"net"

	"github.com/Juniper/contrail-go-api/types"
	log "github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/common"
	"github.com/codilime/contrail-windows-docker/controller"
	"github.com/codilime/contrail-windows-docker/hns"
	dockerTypes "github.com/docker/docker/api/types"
	dockerTypesContainer "github.com/docker/docker/api/types/container"
	dockerTypesNetwork "github.com/docker/docker/api/types/network"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/go-connections/sockets"
	"github.com/docker/go-plugins-helpers/network"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var netAdapter string
var controllerAddr string
var controllerPort int
var useActualController bool

func init() {
	flag.StringVar(&netAdapter, "netAdapter", "Ethernet0",
		"Network adapter to connect HNS switch to")
	flag.StringVar(&controllerAddr, "controllerAddr",
		"10.7.0.54", "Contrail controller addr")
	flag.IntVar(&controllerPort, "controllerPort", 8082, "Contrail controller port")
	flag.BoolVar(&useActualController, "useActualController", true,
		"Whether to use mocked controller or actual.")

	log.SetLevel(log.DebugLevel)
}

func TestDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contrail Network Driver test suite")
}

var _ = BeforeSuite(func() {
	err := common.HardResetHNS()
	Expect(err).ToNot(HaveOccurred())
	err = common.RestartDocker()
	Expect(err).ToNot(HaveOccurred())
})

var contrailController *controller.Controller
var contrailDriver *ContrailDriver
var project *types.Project

const (
	tenantName  = "agatka"
	networkName = "test_net"
	subnetCIDR  = "10.10.10.0/24"
	defaultGW   = "10.10.10.1"
)

var _ = Describe("Contrail Network Driver", func() {

	BeforeEach(func() {
		contrailDriver, contrailController, project = startDriver()
	})

	It("can start and stop listening on a named pipe", func() {
		err := contrailDriver.StartServing()
		Expect(err).ToNot(HaveOccurred())

		d, err := sockets.DialPipe("//./pipe/"+common.DriverName, time.Second*3)
		Expect(err).ToNot(HaveOccurred())
		d.Close()

		err = contrailDriver.StopServing()
		Expect(err).ToNot(HaveOccurred())

		_, err = sockets.DialPipe("//./pipe/"+common.DriverName, time.Second*3)
		Expect(err).To(HaveOccurred())
	})

	It("creates a spec file for duration of listening", func() {
		err := contrailDriver.StartServing()
		Expect(err).ToNot(HaveOccurred())

		_, err = os.Stat(common.PluginSpecFilePath())
		Expect(os.IsNotExist(err)).To(BeFalse())

		err = contrailDriver.StopServing()
		Expect(err).ToNot(HaveOccurred())

		_, err = os.Stat(common.PluginSpecFilePath())
		Expect(os.IsNotExist(err)).To(BeTrue())
	})

	Specify("stopping pipe listener won't cause docker restart to fail", func() {
		err := contrailDriver.StartServing()
		Expect(err).ToNot(HaveOccurred())

		// make sure docker knows about our driver by creating a network
		controller.CreateMockedNetworkWithSubnet(
			contrailController.ApiClient, networkName, subnetCIDR, project)
		_ = createDummyDockerNetwork(tenantName, networkName)

		err = contrailDriver.StopServing()
		Expect(err).ToNot(HaveOccurred())

		err = common.RestartDocker()
		Expect(err).ToNot(HaveOccurred())
	})

	Context("hns and docker are fresh", func() {

		AfterEach(func() {
			err := common.HardResetHNS()
			Expect(err).ToNot(HaveOccurred())
			err = common.RestartDocker()
			Expect(err).ToNot(HaveOccurred())
		})

		Context("on request from docker daemon", func() {

			Context("on GetCapabilities", func() {
				It("returns local scope CapabilitiesResponse, nil", func() {
					resp, err := contrailDriver.GetCapabilities()
					Expect(resp).To(Equal(&network.CapabilitiesResponse{Scope: "local"}))
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("on CreateNetwork", func() {

				var req *network.CreateNetworkRequest
				var genericOptions map[string]interface{}
				BeforeEach(func() {
					req = &network.CreateNetworkRequest{
						NetworkID: "MyAwesomeNet",
						Options:   make(map[string]interface{}),
					}
					genericOptions = make(map[string]interface{})
				})

				type TestCase struct {
					tenant  string
					network string
				}
				DescribeTable("with different, invalid options",
					func(t TestCase) {
						if t.tenant != "" {
							genericOptions["tenant"] = t.tenant
						}
						if t.network != "" {
							genericOptions["network"] = t.network
						}
						req.Options["com.docker.network.generic"] = genericOptions
						err := contrailDriver.CreateNetwork(req)
						Expect(err).To(HaveOccurred())
					},
					Entry("subnet doesn't exist in Contrail", TestCase{
						tenant:  tenantName,
						network: "nonexistingNetwork",
					}),
					Entry("tenant doesn't exist in Contrail", TestCase{
						tenant:  "nonexistingTenant",
						network: networkName,
					}),
					Entry("tenant is not specified in request Options", TestCase{
						network: networkName,
					}),
					Entry("network is not specified in request Options", TestCase{
						tenant: tenantName,
					}),
				)

				Context("tenant and subnet exist in Contrail", func() {
					BeforeEach(func() {
						controller.CreateMockedNetworkWithSubnet(contrailController.ApiClient,
							networkName, subnetCIDR, project)

						genericOptions["network"] = networkName
						genericOptions["tenant"] = tenantName
						req.Options["com.docker.network.generic"] = genericOptions
					})
					It("responds with nil", func() {
						err := contrailDriver.CreateNetwork(req)
						Expect(err).ToNot(HaveOccurred())
					})
					It("creates a HNS network", func() {
						netsBefore, err := hns.ListHNSNetworks()
						Expect(err).ToNot(HaveOccurred())

						err = contrailDriver.CreateNetwork(req)
						Expect(err).ToNot(HaveOccurred())

						netsAfter, err := hns.ListHNSNetworks()
						Expect(err).ToNot(HaveOccurred())
						Expect(netsBefore).To(HaveLen(len(netsAfter) - 1))
					})
				})
			})

			PContext("on AllocateNetwork", func() {
				It("responds with not implemented error", func() {
					req := network.AllocateNetworkRequest{}
					_, err := contrailDriver.AllocateNetwork(&req)
					Expect(err).To(HaveOccurred())
				})
			})

			PContext("on DeleteNetwork", func() {

				Context("docker network has invalid tenant and network name options", func() {})

				Context("docker network exists with valid tenant and net name", func() {
					var dockerNetID string
					var req *network.DeleteNetworkRequest
					BeforeEach(func() {
						dockerNetID = createDummyDockerNetwork(tenantName, networkName)
						req = &network.DeleteNetworkRequest{
							NetworkID: dockerNetID,
						}
					})

					Context("HNS network exists", func() {
						var hnsNetID string
						BeforeEach(func() {
							hnsNetID = hns.MockHNSNetwork(networkName, netAdapter, subnetCIDR,
								defaultGW)
							_, err := hns.GetHNSNetwork(hnsNetID)
							Expect(err).ToNot(HaveOccurred())
						})
						Context("network has no active endpoints", func() {
							It("removes HNS net", func() {
								err := contrailDriver.DeleteNetwork(req)
								Expect(err).ToNot(HaveOccurred())
								_, err = hns.GetHNSNetwork(hnsNetID)
								Expect(err).To(HaveOccurred())
							})
						})

						Context("network has active endpoints", func() {
							BeforeEach(func() {
								hns.MockHNSEndpoint(hnsNetID)
							})
							It("responds with error", func() {
								err := contrailDriver.DeleteNetwork(req)
								Expect(err).To(HaveOccurred())
							})
						})
					})
				})
			})

			PContext("on FreeNetwork", func() {
				It("responds with not implemented error", func() {
					req := network.FreeNetworkRequest{}
					err := contrailDriver.FreeNetwork(&req)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("on CreateEndpoint", func() {

				Context("given a Docker Contrail network", func() {

					dockerNetID := ""
					var docker *dockerClient.Client

					BeforeEach(func() {
						var err error
						docker, err = dockerClient.NewEnvClient()
						Expect(err).ToNot(HaveOccurred())

						err = contrailDriver.StartServing()
						Expect(err).ToNot(HaveOccurred())

						controller.CreateMockedNetworkWithSubnet(
							contrailController.ApiClient, networkName, subnetCIDR, project)

						dockerNetID = createDummyDockerNetwork(tenantName, networkName)
					})

					AfterEach(func() {
						err := contrailDriver.StopServing()
						Expect(err).ToNot(HaveOccurred())

						endpoints, err := hns.ListHNSEndpoints()
						Expect(err).ToNot(HaveOccurred())
						for _, e := range endpoints {
							err = hns.DeleteHNSEndpoint(e.Id)
							Expect(err).ToNot(HaveOccurred())
						}

						endpointsAfter, err := hns.ListHNSEndpoints()
						Expect(err).ToNot(HaveOccurred())
						Expect(endpointsAfter).To(BeEmpty())

						err = docker.NetworkRemove(context.Background(), dockerNetID)
						Expect(err).ToNot(HaveOccurred())
					})

					Context("on Docker container create request", func() {

						var containerID string

						BeforeEach(func() {
							resp, err := docker.ContainerCreate(context.Background(),
								&dockerTypesContainer.Config{
									Image: "microsoft/nanoserver",
								},
								&dockerTypesContainer.HostConfig{
									NetworkMode: networkName,
								},
								nil, "test_container_name")
							Expect(err).ToNot(HaveOccurred())
							containerID = resp.ID
							Expect(containerID).ToNot(Equal(""))

							err = docker.ContainerStart(context.Background(), containerID,
								dockerTypes.ContainerStartOptions{})
							Expect(err).ToNot(HaveOccurred())
						})

						AfterEach(func() {
							timeout := time.Second * 5
							err := docker.ContainerStop(context.Background(), containerID,
								&timeout)
							Expect(err).ToNot(HaveOccurred())

							err = docker.ContainerRemove(context.Background(), containerID,
								dockerTypes.ContainerRemoveOptions{})
							Expect(err).ToNot(HaveOccurred())
						})

						It("allocates Contrail resources", func() {
							net, err := types.VirtualNetworkByName(contrailController.ApiClient,
								fmt.Sprintf("%s:%s:%s", common.DomainName, tenantName,
									networkName))
							Expect(err).ToNot(HaveOccurred())
							Expect(net).ToNot(BeNil())

							// TODO JW-187. For now, VM name is the same as Endpoint ID, not
							// Container ID
							dockerNet, err := docker.NetworkInspect(context.Background(),
								dockerNetID)
							Expect(err).ToNot(HaveOccurred())
							vmName := dockerNet.Containers[containerID].EndpointID

							inst, err := types.VirtualMachineByName(contrailController.ApiClient,
								vmName)
							Expect(err).ToNot(HaveOccurred())
							Expect(inst).ToNot(BeNil())

							vif, err := types.VirtualMachineInterfaceByName(
								contrailController.ApiClient, inst.GetName())
							Expect(err).ToNot(HaveOccurred())
							Expect(vif).ToNot(BeNil())

							ip, err := types.InstanceIpByName(contrailController.ApiClient,
								vif.GetName())
							Expect(err).ToNot(HaveOccurred())
							Expect(ip).ToNot(BeNil())

							ipams, err := net.GetNetworkIpamRefs()
							Expect(err).ToNot(HaveOccurred())
							subnets := ipams[0].Attr.(types.VnSubnetsType).IpamSubnets
							gw := subnets[0].DefaultGateway
							Expect(gw).ToNot(Equal(""))

							macs := vif.GetVirtualMachineInterfaceMacAddresses()
							Expect(macs.MacAddress).To(HaveLen(1))
						})

						It("configures HNS endpoint", func() {
							resp, err := docker.ContainerInspect(context.Background(), containerID)
							Expect(err).ToNot(HaveOccurred())
							Expect(resp).ToNot(BeNil())
							ip := resp.NetworkSettings.Networks[networkName].IPAddress
							mac := resp.NetworkSettings.Networks[networkName].MacAddress
							gw := resp.NetworkSettings.Networks[networkName].Gateway

							endpoints, err := hns.ListHNSEndpoints()
							Expect(err).ToNot(HaveOccurred())
							Expect(endpoints).ToNot(BeEmpty())

							ep := endpoints[0]
							Expect(ep).ToNot(BeNil(), "Endpoint not found")
							Expect(ep.IPAddress).To(Equal(net.ParseIP(ip)))
							formattedMac := strings.Replace(strings.ToUpper(mac), ":", "-", -1)
							Expect(ep.MacAddress).To(Equal(formattedMac))
							Expect(ep.GatewayAddress).To(Equal(gw))
						})

						PIt("configures vRouter agent", func() {})
					})

				})

				Context("docker network doesn't exist", func() {
					It("responds with err", func() {
						req := &network.CreateEndpointRequest{
							EndpointID: "somecontainerID",
						}
						_, err := contrailDriver.CreateEndpoint(req)
						Expect(err).To(HaveOccurred())
					})
				})

				PContext("docker network is misconfigured", func() {
					It("responds with err", func() {
						docker, err := dockerClient.NewEnvClient()
						Expect(err).ToNot(HaveOccurred())
						params := &dockerTypes.NetworkCreate{
							Options: map[string]string{
								"tenant":  tenantName,
								"network": "lolel",
							},
						}
						_, err = docker.NetworkCreate(context.Background(), networkName, *params)
						Expect(err).ToNot(HaveOccurred())
						req := network.CreateEndpointRequest{
							EndpointID: "somecontainerID",
						}
						resp, err := contrailDriver.CreateEndpoint(&req)
						Expect(err).To(HaveOccurred())
						Expect(resp).To(BeNil())
					})
				})
			})

			Context("on DeleteEndpoint", func() {
				PIt("deallocates Contrail resources", func() {})
				PIt("configures vRouter Agent", func() {})
				It("responds with nil", func() {
					req := network.DeleteEndpointRequest{}
					err := contrailDriver.DeleteEndpoint(&req)
					Expect(err).ToNot(HaveOccurred())
				})
			})

			PContext("on EndpointInfo", func() {
				It("responds with proper InfoResponse", func() {})
			})

			PContext("on Join", func() {
				Context("queried endpoint exists", func() {
					It("responds with proper JoinResponse", func() {}) // nil maybe?
				})

				Context("queried endpoint doesn't exist", func() {
					It("responds with err", func() {})
				})
			})

			PContext("on Leave", func() {

				Context("queried endpoint exists", func() {
					It("responds with proper JoinResponse, nil", func() {})
				})

				Context("queried endpoint doesn't exist", func() {
					It("responds with err", func() {})
				})
			})

			Context("on DiscoverNew", func() {
				It("responds with nil", func() {
					req := network.DiscoveryNotification{}
					err := contrailDriver.DiscoverNew(&req)
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("on DiscoverDelete", func() {
				It("responds with nil", func() {
					req := network.DiscoveryNotification{}
					err := contrailDriver.DiscoverDelete(&req)
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("on ProgramExternalConnectivity", func() {
				It("responds with nil", func() {
					req := network.ProgramExternalConnectivityRequest{}
					err := contrailDriver.ProgramExternalConnectivity(&req)
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("on RevokeExternalConnectivity", func() {
				It("responds with nil", func() {
					req := network.RevokeExternalConnectivityRequest{}
					err := contrailDriver.RevokeExternalConnectivity(&req)
					Expect(err).ToNot(HaveOccurred())
				})
			})

		})
	})

})

func startDriver() (*ContrailDriver, *controller.Controller, *types.Project) {
	var c *controller.Controller
	var p *types.Project

	if useActualController {
		c, p = controller.NewClientAndProject(tenantName, controllerAddr, controllerPort)
	} else {
		c, p = controller.NewMockedClientAndProject(tenantName)
	}
	d, err := NewDriver("Ethernet0", c)
	Expect(err).ToNot(HaveOccurred())

	return d, c, p
}

func createDummyDockerNetwork(tenant, netName string) string {
	docker, err := dockerClient.NewEnvClient()
	Expect(err).ToNot(HaveOccurred())
	params := &dockerTypes.NetworkCreate{
		Driver: common.DriverName,
		IPAM: &dockerTypesNetwork.IPAM{
			// libnetwork/ipams/windowsipam ("windows") driver is a null ipam driver.
			// We use 0/32 subnet because no preferred address is specified (as documented in
			// source code of windowsipam driver). We do this because our driver has to handle
			// IP assignment.
			// If container has IP before CreateEndpoint request is handled and CreateEndpoint
			// returns a new IP (assigned by Contrail), docker daemon will complain that we cannot
			// reassign IPs. Hence, we tell the IPAM driver to not assign any IPs.
			Driver: "windows",
			Config: []dockerTypesNetwork.IPAMConfig{
				{
					Subnet: "0.0.0.0/32",
				},
			},
		},
		Options: map[string]string{
			"tenant":  tenant,
			"network": networkName,
		},
	}
	resp, err := docker.NetworkCreate(context.Background(), networkName, *params)
	Expect(err).ToNot(HaveOccurred())
	return resp.ID
}
