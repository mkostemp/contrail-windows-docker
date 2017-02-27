package driver

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/contrail-go-api/types"
	"github.com/Microsoft/hcsshim"
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

	docker := getDockerClient()
	cleanupAllDockerNetworksAndContainers(docker)
})

var contrailController *controller.Controller
var contrailDriver *ContrailDriver
var project *types.Project

const (
	tenantName  = "agatka"
	networkName = "test_net"
	subnetCIDR  = "10.10.10.0/24"
	defaultGW   = "10.10.10.1"
	timeout     = time.Second * 5
)

var _ = Describe("Contrail Network Driver", func() {

	BeforeEach(func() {
		contrailDriver, contrailController, project = startDriver()
	})

	It("can start and stop listening on a named pipe", func() {
		err := contrailDriver.StartServing()
		Expect(err).ToNot(HaveOccurred())

		d, err := sockets.DialPipe("//./pipe/"+common.DriverName, timeout)
		Expect(err).ToNot(HaveOccurred())
		d.Close()

		err = contrailDriver.StopServing()
		Expect(err).ToNot(HaveOccurred())

		_, err = sockets.DialPipe("//./pipe/"+common.DriverName, timeout)
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
		_ = createContrailNetwork(contrailController)
		docker := getDockerClient()
		_ = createValidDockerNetwork(docker)

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
					ipamData := []*network.IPAMData{
						{
							Pool: subnetCIDR,
						},
					}
					req = &network.CreateNetworkRequest{
						NetworkID: "MyAwesomeNet",
						Options:   make(map[string]interface{}),
						IPv4Data:  ipamData,
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
						_ = createContrailNetwork(contrailController)

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

			Context("on AllocateNetwork", func() {
				It("responds with not implemented error", func() {
					req := network.AllocateNetworkRequest{}
					_, err := contrailDriver.AllocateNetwork(&req)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("on DeleteNetwork", func() {

				dockerNetID := ""
				var docker *dockerClient.Client
				var contrailNet *types.VirtualNetwork

				BeforeEach(func() {
					err := contrailDriver.StartServing()
					Expect(err).ToNot(HaveOccurred())

					contrailNet = createContrailNetwork(contrailController)
					docker = getDockerClient()
					dockerNetID = createValidDockerNetwork(docker)
				})

				AfterEach(func() {
					err := contrailDriver.StopServing()
					Expect(err).ToNot(HaveOccurred())
				})

				assertRemovesHNSNet := func() {
					resp, err := contrailDriver.hnsMgr.GetNetwork(tenantName, networkName,
						subnetCIDR)
					Expect(err).To(HaveOccurred())
					Expect(resp).To(BeNil())
				}

				assertRemovesDockerNet := func() {
					_, err := docker.NetworkInspect(context.Background(), dockerNetID)
					Expect(err).To(HaveOccurred())
				}

				assertDoesNotRemoveContrailNet := func() {
					net, err := types.VirtualNetworkByName(contrailController.ApiClient,
						fmt.Sprintf("%s:%s:%s", common.DomainName, tenantName,
							networkName))
					Expect(err).ToNot(HaveOccurred())
					Expect(net).ToNot(BeNil())
				}

				Context("docker network, Contrail network and HNS network are fine", func() {
					BeforeEach(func() {
						err := removeDockerNetwork(docker, dockerNetID)
						Expect(err).ToNot(HaveOccurred())
					})
					It("removes HNS net", assertRemovesHNSNet)
					It("removes docker net", assertRemovesDockerNet)
					It("does not remove Contrail net", assertDoesNotRemoveContrailNet)
				})

				Context("HNS network doesn't exist", func() {
					// for example, HNS was hard-reset while docker wasn't.
					BeforeEach(func() {
						contrailDriver.hnsMgr.DeleteNetwork(tenantName, networkName, subnetCIDR)
						err := removeDockerNetwork(docker, dockerNetID)
						Expect(err).ToNot(HaveOccurred())
					})
					It("removes docker net", assertRemovesDockerNet)
					It("does not remove Contrail net", assertDoesNotRemoveContrailNet)
				})

				Context("Contrail network doesn't exist", func() {
					// for example, somebody deleted Contrail network before removing docker/hns
					BeforeEach(func() {
						err := contrailController.DeleteElementRecursive(contrailNet)
						Expect(err).ToNot(HaveOccurred())
						err = removeDockerNetwork(docker, dockerNetID)
						Expect(err).ToNot(HaveOccurred())
					})
					It("removes HNS net", assertRemovesHNSNet)
					It("removes docker net", assertRemovesDockerNet)
				})
			})

			Context("on FreeNetwork", func() {
				It("responds with not implemented error", func() {
					req := network.FreeNetworkRequest{}
					err := contrailDriver.FreeNetwork(&req)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("on CreateEndpoint", func() {

				BeforeEach(func() {
					err := contrailDriver.StartServing()
					Expect(err).ToNot(HaveOccurred())
				})

				AfterEach(func() {
					err := contrailDriver.StopServing()
					Expect(err).ToNot(HaveOccurred())
				})

				Context("given a Docker and Contrail networks", func() {

					dockerNetID := ""
					var docker *dockerClient.Client

					BeforeEach(func() {
						_ = createContrailNetwork(contrailController)

						docker = getDockerClient()
						dockerNetID = createValidDockerNetwork(docker)
					})

					AfterEach(func() {
						endpoints, err := hns.ListHNSEndpoints()
						Expect(err).ToNot(HaveOccurred())
						for _, e := range endpoints {
							err = hns.DeleteHNSEndpoint(e.Id)
							Expect(err).ToNot(HaveOccurred())
						}

						endpointsAfter, err := hns.ListHNSEndpoints()
						Expect(err).ToNot(HaveOccurred())
						Expect(endpointsAfter).To(BeEmpty())

						err = removeDockerNetwork(docker, dockerNetID)
						Expect(err).ToNot(HaveOccurred())
					})

					Context("corresponding HNS network doesn't exist", func() {
						// for example, HNS was hard-reset while docker wasn't.
						containerID := ""
						BeforeEach(func() {
							contrailDriver.hnsMgr.DeleteNetwork(tenantName, networkName,
								subnetCIDR)
						})
						AfterEach(func() {
							err := docker.ContainerRemove(context.Background(), containerID,
								dockerTypes.ContainerRemoveOptions{})
							Expect(err).ToNot(HaveOccurred())
						})
						It("responds with err", func() {
							var err error
							containerID, err = runDockerContainer(docker)
							Expect(err).To(HaveOccurred())
						})
					})

					Context("corresponding HNS network exists", func() {

						var containerID string

						BeforeEach(func() {
							var err error
							containerID, err = runDockerContainer(docker)
							Expect(err).ToNot(HaveOccurred())
						})

						AfterEach(func() {
							stopAndRemoveDockerContainer(docker, containerID)
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
							resp, err := docker.ContainerInspect(context.Background(),
								containerID)
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
							Expect(ep.IPAddress.String()).To(Equal(ip))
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

			})

			Context("on DeleteEndpoint", func() {

				dockerNetID := ""
				containerID := ""
				hnsEndpointID := ""
				var docker *dockerClient.Client
				var contrailNet *types.VirtualNetwork
				var contrailInst *types.VirtualMachine
				var contrailVif *types.VirtualMachineInterface
				var contrailIP *types.InstanceIp

				BeforeEach(func() {
					err := contrailDriver.StartServing()
					Expect(err).ToNot(HaveOccurred())

					contrailNet = createContrailNetwork(contrailController)
					docker = getDockerClient()
					dockerNetID = createValidDockerNetwork(docker)

					containerID, err = runDockerContainer(docker)
					Expect(err).ToNot(HaveOccurred())

					hnsNets, err := contrailDriver.hnsMgr.ListNetworks()
					Expect(err).ToNot(HaveOccurred())
					Expect(hnsNets).To(HaveLen(1))
					eps, err := hns.ListHNSEndpointsOfNetwork(hnsNets[0].Id)
					Expect(err).ToNot(HaveOccurred())

					hnsEndpointID = eps[0].Id

					// TODO JW-187. For now, VM name is the same as Endpoint ID, not
					// Container ID
					dockerNet, err := docker.NetworkInspect(context.Background(),
						dockerNetID)
					Expect(err).ToNot(HaveOccurred())
					vmName := dockerNet.Containers[containerID].EndpointID

					contrailInst, err = types.VirtualMachineByName(contrailController.ApiClient,
						vmName)
					Expect(err).ToNot(HaveOccurred())
					Expect(contrailInst).ToNot(BeNil())

					contrailVif, err = types.VirtualMachineInterfaceByName(
						contrailController.ApiClient, contrailInst.GetName())
					Expect(err).ToNot(HaveOccurred())
					Expect(contrailVif).ToNot(BeNil())

					contrailIP, err = types.InstanceIpByName(contrailController.ApiClient,
						contrailVif.GetName())
					Expect(err).ToNot(HaveOccurred())
					Expect(contrailIP).ToNot(BeNil())
				})

				AfterEach(func() {
					cleanupAllDockerNetworksAndContainers(docker)
					err := contrailDriver.StopServing()
					Expect(err).ToNot(HaveOccurred())
				})

				assertRemovesDockerEndpoint := func() {
					_, err := docker.ContainerInspect(context.Background(), containerID)
					Expect(err).To(HaveOccurred())
				}

				assertRemovesHNSEndpoint := func() {
					ep, err := hns.GetHNSEndpoint(hnsEndpointID)
					Expect(err).To(HaveOccurred())
					Expect(ep).To(BeNil())
				}

				assertRemovesContrailVM := func() {
					_, err := types.VirtualMachineByName(contrailController.ApiClient,
						contrailInst.GetName())
					Expect(err).To(HaveOccurred())

					_, err = types.VirtualMachineInterfaceByName(
						contrailController.ApiClient, contrailInst.GetName())
					Expect(err).To(HaveOccurred())

					_, err = types.InstanceIpByName(contrailController.ApiClient,
						contrailVif.GetName())
					Expect(err).To(HaveOccurred())
				}

				Context("happy case: HNS, docker and Contrail states are in sync", func() {
					BeforeEach(func() {
						stopAndRemoveDockerContainer(docker, containerID)
					})
					It("removes docker endpoint", assertRemovesDockerEndpoint)
					It("removes HNS endpoint", assertRemovesHNSEndpoint)
					It("removes virtual-machine and its children in Contrail",
						assertRemovesContrailVM)
					PIt("removes port from vRouter Agent", func() {})
				})

				Context("HNS endpoint doesn't exist", func() {
					BeforeEach(func() {
						err := hns.DeleteHNSEndpoint(hnsEndpointID)
						Expect(err).ToNot(HaveOccurred())
						stopAndRemoveDockerContainer(docker, containerID)
					})
					It("removes docker endpoint", assertRemovesDockerEndpoint)
					It("removes virtual-machine and its children in Contrail",
						assertRemovesContrailVM)
					PIt("removes port from vRouter Agent", func() {})
				})

				Context("virtual-machine in Contrail doesn't exist", func() {
					BeforeEach(func() {
						err := contrailController.DeleteElementRecursive(contrailInst)
						Expect(err).ToNot(HaveOccurred())
						stopAndRemoveDockerContainer(docker, containerID)
					})
					It("removes docker endpoint", assertRemovesDockerEndpoint)
					It("removes HNS endpoint", assertRemovesHNSEndpoint)
					PIt("removes port from vRouter Agent", func() {})
				})

				PContext("port doesn't exist in vRouter Agent", func() {

				})

			})

			Context("docker, HNS and Contrail are setup correctly", func() {

				dockerNetID := ""
				containerID := ""
				var docker *dockerClient.Client
				var contrailNet *types.VirtualNetwork
				var contrailGW string

				deleteQueriedEndpoint := func() {
					hnsNets, err := contrailDriver.hnsMgr.ListNetworks()
					Expect(err).ToNot(HaveOccurred())
					eps, err := hns.ListHNSEndpointsOfNetwork(hnsNets[0].Id)
					Expect(err).ToNot(HaveOccurred())
					hnsEndpointID := eps[0].Id
					err = hns.DeleteHNSEndpoint(hnsEndpointID)
					Expect(err).ToNot(HaveOccurred())
				}

				BeforeEach(func() {
					err := contrailDriver.StartServing()
					Expect(err).ToNot(HaveOccurred())

					contrailNet = createContrailNetwork(contrailController)
					ipams, err := contrailNet.GetNetworkIpamRefs()
					Expect(err).ToNot(HaveOccurred())
					subnets := ipams[0].Attr.(types.VnSubnetsType).IpamSubnets
					contrailGW = subnets[0].DefaultGateway

					docker = getDockerClient()
					dockerNetID = createValidDockerNetwork(docker)
					containerID, err = runDockerContainer(docker)
					Expect(err).ToNot(HaveOccurred())
				})

				AfterEach(func() {
					cleanupAllDockerNetworksAndContainers(docker)
					err := contrailDriver.StopServing()
					Expect(err).ToNot(HaveOccurred())
				})

				Context("on EndpointInfo", func() {
					var req *network.InfoRequest

					BeforeEach(func() {
						dockerNet, err := docker.NetworkInspect(context.Background(), dockerNetID)
						Expect(err).ToNot(HaveOccurred())
						req = &network.InfoRequest{
							NetworkID:  dockerNetID,
							EndpointID: dockerNet.Containers[containerID].EndpointID,
						}
					})

					Context("queried endpoint exists", func() {
						var hnsEndpoint *hcsshim.HNSEndpoint
						BeforeEach(func() {
							hnsNets, err := contrailDriver.hnsMgr.ListNetworks()
							Expect(err).ToNot(HaveOccurred())
							eps, err := hns.ListHNSEndpointsOfNetwork(hnsNets[0].Id)
							Expect(err).ToNot(HaveOccurred())
							hnsEndpointID := eps[0].Id
							hnsEndpoint, err = hns.GetHNSEndpoint(hnsEndpointID)
							Expect(err).ToNot(HaveOccurred())

						})
						It("responds with proper InfoResponse", func() {
							resp, err := contrailDriver.EndpointInfo(req)
							Expect(err).ToNot(HaveOccurred())
							Expect(resp.Value).To(HaveKeyWithValue("hnsid", hnsEndpoint.Id))
							Expect(resp.Value).To(HaveKeyWithValue(
								"com.docker.network.endpoint.macaddress", hnsEndpoint.MacAddress))
						})
					})

					Context("queried endpoint doesn't exist", func() {
						BeforeEach(deleteQueriedEndpoint)
						It("responds with err", func() {
							_, err := contrailDriver.EndpointInfo(req)
							Expect(err).To(HaveOccurred())
						})
					})
				})

				Context("on Join", func() {
					var req *network.JoinRequest

					BeforeEach(func() {
						dockerNet, err := docker.NetworkInspect(context.Background(), dockerNetID)
						Expect(err).ToNot(HaveOccurred())
						req = &network.JoinRequest{
							NetworkID:  dockerNetID,
							EndpointID: dockerNet.Containers[containerID].EndpointID,
						}
					})

					Context("queried endpoint exists", func() {
						It("responds with valid gateway IP and disabled gw service", func() {
							resp, err := contrailDriver.Join(req)
							Expect(err).ToNot(HaveOccurred())
							Expect(resp.DisableGatewayService).To(BeTrue())
							Expect(resp.Gateway).To(Equal(contrailGW))
						})
					})

					Context("queried endpoint doesn't exist", func() {
						BeforeEach(deleteQueriedEndpoint)
						It("responds with err", func() {
							_, err := contrailDriver.Join(req)
							Expect(err).To(HaveOccurred())
						})
					})
				})

				Context("on Leave", func() {
					var req *network.LeaveRequest

					BeforeEach(func() {
						dockerNet, err := docker.NetworkInspect(context.Background(), dockerNetID)
						Expect(err).ToNot(HaveOccurred())
						req = &network.LeaveRequest{
							NetworkID:  dockerNetID,
							EndpointID: dockerNet.Containers[containerID].EndpointID,
						}
					})

					Context("queried endpoint exists", func() {
						It("responds with nil", func() {
							err := contrailDriver.Leave(req)
							Expect(err).ToNot(HaveOccurred())
						})
					})

					Context("queried endpoint doesn't exist", func() {
						BeforeEach(deleteQueriedEndpoint)
						It("responds with err", func() {
							err := contrailDriver.Leave(req)
							Expect(err).To(HaveOccurred())
						})
					})
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
	d := NewDriver(netAdapter, c)

	return d, c, p
}

func getDockerClient() *dockerClient.Client {
	docker, err := dockerClient.NewEnvClient()
	Expect(err).ToNot(HaveOccurred())
	return docker
}

func runDockerContainer(docker *dockerClient.Client) (string, error) {
	resp, err := docker.ContainerCreate(context.Background(),
		&dockerTypesContainer.Config{
			Image: "microsoft/nanoserver",
		},
		&dockerTypesContainer.HostConfig{
			NetworkMode: networkName,
		},
		nil, "test_container_name")
	Expect(err).ToNot(HaveOccurred())
	containerID := resp.ID
	Expect(containerID).ToNot(Equal(""))

	err = docker.ContainerStart(context.Background(), containerID,
		dockerTypes.ContainerStartOptions{})

	return containerID, err
}

func stopAndRemoveDockerContainer(docker *dockerClient.Client, containerID string) {
	timeout := time.Second * 5
	err := docker.ContainerStop(context.Background(), containerID,
		&timeout)
	Expect(err).ToNot(HaveOccurred())

	err = docker.ContainerRemove(context.Background(), containerID,
		dockerTypes.ContainerRemoveOptions{Force: true})
	Expect(err).ToNot(HaveOccurred())
}

func createValidDockerNetwork(docker *dockerClient.Client) string {
	return createDockerNetwork(tenantName, networkName, docker)
}

func createDockerNetwork(tenant, network string, docker *dockerClient.Client) string {
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
			"network": network,
		},
	}
	resp, err := docker.NetworkCreate(context.Background(), networkName, *params)
	Expect(err).ToNot(HaveOccurred())
	return resp.ID
}

func removeDockerNetwork(docker *dockerClient.Client, dockerNetID string) error {
	return docker.NetworkRemove(context.Background(), dockerNetID)
}

func cleanupAllDockerNetworksAndContainers(docker *dockerClient.Client) {
	nets, err := docker.NetworkList(context.Background(), dockerTypes.NetworkListOptions{})
	Expect(err).ToNot(HaveOccurred())
	for _, net := range nets {
		log.Debugln("Cleaning up endpoints of docker network", net.Name)
		for containerID, _ := range net.Containers {
			log.Debugln("Stopping and removing container", containerID)
			stopAndRemoveDockerContainer(docker, containerID)
		}
		if net.Name == "none" || net.Name == "nat" {
			// those networks are pre-defined and cannot be removed (will cause error)
			continue
		}
		log.Debugln("Removing docker network", net.Name)
		err = removeDockerNetwork(docker, net.ID)
		Expect(err).ToNot(HaveOccurred())
	}
}

func createContrailNetwork(c *controller.Controller) *types.VirtualNetwork {
	return controller.CreateMockedNetworkWithSubnet(
		c.ApiClient, networkName, subnetCIDR, project)
}
