package controller

import (
	"flag"
	"testing"

	contrail "github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	log "github.com/Sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var controllerAddr string
var controllerPort int
var useActualController bool

func init() {
	flag.StringVar(&controllerAddr, "controllerAddr",
		"10.7.0.54", "Contrail controller addr")
	flag.IntVar(&controllerPort, "controllerPort", 8082, "Contrail controller port")
	flag.BoolVar(&useActualController, "useActualController", true,
		"Whether to use mocked controller or actual.")

	log.SetLevel(log.DebugLevel)
}

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller client test suite")
}

const (
	tenantName   = "agatka"
	networkName  = "test_net"
	subnetCIDR   = "10.10.10.0/24"
	subnetPrefix = "10.10.10.0"
	subnetMask   = 24
	defaultGW    = "10.10.10.1"
	ifaceMac     = "contrail_pls_check_macs"
	containerID  = "12345678901"
)

var _ = BeforeSuite(func() {
	if useActualController {
		// this cleans up
		client, _ := NewClientAndProject(tenantName, controllerAddr, controllerPort)
		CleanupLingeringVM(client, containerID)
	}
})

var _ = Describe("Controller", func() {

	var client *Controller
	var project *types.Project

	BeforeEach(func() {
		if useActualController {
			client, project = NewClientAndProject(tenantName, controllerAddr, controllerPort)
		} else {
			client, project = NewMockedClientAndProject(tenantName)
		}
	})

	AfterEach(func() {
		if useActualController {
			CleanupLingeringVM(client, containerID)
		}
	})

	Specify("cleaning up resources that are referred to by two other doesn't fail", func() {
		// instanceIP and VMI are both referred to by virtual network, and instanceIP refers
		// to VMI
		testNetwork := CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
			project)
		testInstance := CreateMockedInstance(client.ApiClient, tenantName, containerID)
		testInterface := CreateMockedInterface(client.ApiClient, testInstance, testNetwork)
		_ = CreateMockedInstanceIP(client.ApiClient, tenantName, testInterface,
			testNetwork)

		// shouldn't error when creating new client and project
		if useActualController {
			client, project = NewClientAndProject(tenantName, controllerAddr, controllerPort)
		} else {
			client, project = NewMockedClientAndProject(tenantName)
		}
	})

	Specify("recursive deletion removes elements down the ref tree", func() {
		testNetwork := CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
			project)
		testInstance := CreateMockedInstance(client.ApiClient, tenantName, containerID)
		testInterface := CreateMockedInterface(client.ApiClient, testInstance, testNetwork)
		testInstanceIP := CreateMockedInstanceIP(client.ApiClient, tenantName, testInterface,
			testNetwork)

		err := client.DeleteElementRecursive(testInstance)
		Expect(err).ToNot(HaveOccurred())

		_, err = client.ApiClient.FindByUuid(testNetwork.GetType(), testNetwork.GetUuid())
		Expect(err).ToNot(HaveOccurred())

		for _, supposedToBeRemovedObject := range []contrail.IObject{testInstance, testInterface,
			testInstanceIP} {
			_, err = client.ApiClient.FindByUuid(supposedToBeRemovedObject.GetType(),
				supposedToBeRemovedObject.GetUuid())
			Expect(err).To(HaveOccurred())
		}
	})

	Describe("getting Contrail network", func() {
		Context("when network already exists in Contrail", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName,
					subnetCIDR, project)
			})
			It("returns it", func() {
				net, err := client.GetNetwork(tenantName, networkName)
				Expect(err).ToNot(HaveOccurred())
				Expect(net.GetUuid()).To(Equal(testNetwork.GetUuid()))
			})
		})
		Context("when network doesn't exist in Contrail", func() {
			It("returns an error", func() {
				net, err := client.GetNetwork(tenantName, networkName)
				Expect(err).To(HaveOccurred())
				Expect(net).To(BeNil())
			})
		})
	})

	Describe("getting Contrail subnet info", func() {
		assertGettingSubnetFails := func(getTestedNet func() *types.VirtualNetwork,
			CIDR string) func() {
			return func() {
				_, err := client.GetIpamSubnet(getTestedNet(), CIDR)
				Expect(err).To(HaveOccurred())
			}
		}
		Context("network has one subnet with default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
				AddSubnetWithDefaultGateway(client.ApiClient, subnetPrefix, defaultGW,
					subnetMask, testNetwork)
			})
			Specify("getting subnet meta works", func() {
				ipam, err := client.GetIpamSubnet(testNetwork, "")
				Expect(err).ToNot(HaveOccurred())
				Expect(ipam.DefaultGateway).To(Equal(defaultGW))
				Expect(ipam.Subnet.IpPrefix).To(Equal(subnetPrefix))
				Expect(ipam.Subnet.IpPrefixLen).To(Equal(subnetMask))
			})
			Specify("getting subnet when specifying CIDR works", func() {
				_, err := client.GetIpamSubnet(testNetwork, subnetCIDR)
				Expect(err).ToNot(HaveOccurred())
			})
			Specify("getting subnet when specifying CIDR not in Contrail fails",
				assertGettingSubnetFails(func() *types.VirtualNetwork {
					return testNetwork
				}, "1.2.3.4/16"))
		})
		Context("network has one subnet without default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName,
					subnetCIDR, project)
			})
			Specify("getting default gw IP returns error", func() {
				ipam, err := client.GetIpamSubnet(testNetwork, "")
				Expect(err).ToNot(HaveOccurred())
				if useActualController {
					Expect(ipam.DefaultGateway).ToNot(Equal(""))
					Expect(err).ToNot(HaveOccurred())
				} else {
					// mocked controller lacks some logic here
					Expect(ipam.DefaultGateway).To(Equal(""))
					Expect(err).To(HaveOccurred())
				}
			})
			Specify("getting subnet prefix and prefix len works", func() {
				ipam, err := client.GetIpamSubnet(testNetwork, "")
				Expect(err).ToNot(HaveOccurred())
				Expect(ipam.Subnet.IpPrefix).To(Equal(subnetPrefix))
				Expect(ipam.Subnet.IpPrefixLen).To(Equal(subnetMask))
			})
		})
		Context("network doesn't have subnets", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
			})
			Specify("getting subnet returns error",
				assertGettingSubnetFails(func() *types.VirtualNetwork {
					return testNetwork
				}, ""))
		})
		Context("network has multiple subnets", func() {
			var testNetwork *types.VirtualNetwork
			const (
				prefix1 = "10.10.10.0"
				gw1     = "10.10.10.1"
				cidr1   = "10.10.10.0/24"
				prefix2 = "10.20.20.0"
				gw2     = "10.20.20.1"
				cidr2   = "10.20.20.0/24"
			)
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
				AddSubnetWithDefaultGateway(client.ApiClient, prefix1, gw1, 24,
					testNetwork)
				AddSubnetWithDefaultGateway(client.ApiClient, prefix2, gw2, 24,
					testNetwork)
			})
			Context("user specified valid subnet", func() {
				Specify("getting specific subnets works", func() {
					ipam1, err := client.GetIpamSubnet(testNetwork, cidr1)
					Expect(err).ToNot(HaveOccurred())
					Expect(ipam1.DefaultGateway).To(Equal(gw1))

					ipam2, err := client.GetIpamSubnet(testNetwork, cidr2)
					Expect(err).ToNot(HaveOccurred())
					Expect(ipam2.DefaultGateway).To(Equal(gw2))
				})
			})
			Context("user didn't specify a subnet", func() {
				Specify("getting subnet1 returns error",
					assertGettingSubnetFails(func() *types.VirtualNetwork {
						return testNetwork
					}, ""))
				Specify("getting subnet2 returns error",
					assertGettingSubnetFails(func() *types.VirtualNetwork {
						return testNetwork
					}, ""))
			})
			Context("user specified invalid subnet", func() {
				Specify("getting subnet1 returns error",
					assertGettingSubnetFails(func() *types.VirtualNetwork {
						return testNetwork
					}, "10.12.13.0/24"))
				Specify("getting subnet2 returns error",
					assertGettingSubnetFails(func() *types.VirtualNetwork {
						return testNetwork
					}, "10.12.13.0/24"))
			})
		})
	})

	Describe("getting Contrail instance", func() {
		Context("when instance already exists in Contrail", func() {
			var testInstance *types.VirtualMachine
			BeforeEach(func() {
				testInstance = CreateMockedInstance(client.ApiClient, tenantName, containerID)
			})
			It("returns existing instance", func() {
				instance, err := client.GetOrCreateInstance(tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(instance).ToNot(BeNil())
				Expect(instance.GetUuid()).To(Equal(testInstance.GetUuid()))
			})
		})
		Context("when instance doesn't exist in Contrail", func() {
			It("creates a new instance", func() {
				instance, err := client.GetOrCreateInstance(tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(instance).ToNot(BeNil())

				existingInst, err := types.VirtualMachineByUuid(client.ApiClient,
					instance.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingInst.GetUuid()).To(Equal(instance.GetUuid()))
			})
		})
	})

	Describe("getting Contrail virtual interface", func() {
		var testNetwork *types.VirtualNetwork
		var testInstance *types.VirtualMachine
		BeforeEach(func() {
			testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
			testInstance = CreateMockedInstance(client.ApiClient, tenantName, containerID)
		})
		Context("when vif already exists in Contrail", func() {
			var testInterface *types.VirtualMachineInterface
			BeforeEach(func() {
				testInterface = CreateMockedInterface(client.ApiClient, testInstance,
					testNetwork)
			})
			It("returns existing vif", func() {
				iface, err := client.GetOrCreateInterface(testNetwork, testInstance)
				Expect(err).ToNot(HaveOccurred())
				Expect(iface).ToNot(BeNil())
				Expect(iface.GetUuid()).To(Equal(testInterface.GetUuid()))
			})
		})
		Context("when vif doesn't exist in Contrail", func() {
			It("creates a new vif", func() {
				iface, err := client.GetOrCreateInterface(testNetwork, testInstance)
				Expect(err).ToNot(HaveOccurred())
				Expect(iface).ToNot(BeNil())

				existingIface, err := types.VirtualMachineInterfaceByUuid(client.ApiClient,
					iface.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingIface.GetUuid()).To(Equal(iface.GetUuid()))
			})
		})
	})

	Describe("getting virtual interface MAC", func() {
		var testNetwork *types.VirtualNetwork
		var testInstance *types.VirtualMachine
		var testInterface *types.VirtualMachineInterface
		BeforeEach(func() {
			testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
			testInstance = CreateMockedInstance(client.ApiClient, tenantName, containerID)
			testInterface = CreateMockedInterface(client.ApiClient, testInstance,
				testNetwork)
		})
		Context("when vif has MAC", func() {
			BeforeEach(func() {
				AddMacToInterface(client.ApiClient, ifaceMac, testInterface)
			})
			It("returns MAC address", func() {
				mac, err := client.GetInterfaceMac(testInterface)
				Expect(err).ToNot(HaveOccurred())
				Expect(mac).To(Equal(ifaceMac))
			})
		})
		Context("when vif doesn't have a MAC", func() {
			It("returns error", func() {
				mac, err := client.GetInterfaceMac(testInterface)
				Expect(err).To(HaveOccurred())
				Expect(mac).To(Equal(""))
			})
		})
	})

	Describe("getting Contrail instance IP", func() {
		var testNetwork *types.VirtualNetwork
		var testInstance *types.VirtualMachine
		var testInterface *types.VirtualMachineInterface
		BeforeEach(func() {
			testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
			testInstance = CreateMockedInstance(client.ApiClient, tenantName, containerID)
			testInterface = CreateMockedInterface(client.ApiClient, testInstance,
				testNetwork)
		})
		Context("when instance IP already exists in Contrail", func() {
			var testInstanceIP *types.InstanceIp
			BeforeEach(func() {
				testInstanceIP = CreateMockedInstanceIP(client.ApiClient, tenantName,
					testInterface, testNetwork)
			})
			It("returns existing instance IP", func() {
				instanceIP, err := client.GetOrCreateInstanceIp(testNetwork, testInterface)
				Expect(err).ToNot(HaveOccurred())
				Expect(instanceIP).ToNot(BeNil())
				Expect(instanceIP.GetUuid()).To(Equal(testInstanceIP.GetUuid()))
				Expect(instanceIP.GetInstanceIpAddress()).To(Equal(
					testInstanceIP.GetInstanceIpAddress()))

				existingIP, err := types.InstanceIpByUuid(client.ApiClient, instanceIP.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingIP.GetUuid()).To(Equal(instanceIP.GetUuid()))
			})
		})
		Context("when instance IP doesn't exist in Contrail", func() {
			It("creates new instance IP", func() {
				instanceIP, err := client.GetOrCreateInstanceIp(testNetwork, testInterface)
				Expect(err).ToNot(HaveOccurred())
				Expect(instanceIP).ToNot(BeNil())
				Expect(instanceIP.GetInstanceIpAddress()).ToNot(Equal(""))

				existingIP, err := types.InstanceIpByUuid(client.ApiClient, instanceIP.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingIP.GetUuid()).To(Equal(instanceIP.GetUuid()))
			})
		})
	})
})

var _ = Describe("Authenticating", func() {

	type TestCase struct {
		shouldErr bool
		keys      KeystoneEnvs
	}
	DescribeTable("with different keystone env variables",
		func(t TestCase) {
			_, err := NewController(controllerAddr, controllerPort, &t.keys)
			if t.shouldErr {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(err).ToNot(HaveOccurred())
			}
		},
		Entry("env variables are not set", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "",
				os_username:    "",
				os_tenant_name: "",
				os_password:    "",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("bad url", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "http://10.7.0.54:5000/",
				os_username:    "admin",
				os_tenant_name: "admin",
				os_password:    "secret123",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("empty url", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "",
				os_username:    "admin",
				os_tenant_name: "admin",
				os_password:    "secret123",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("bad user", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "http://10.7.0.54:5000/v2.0",
				os_username:    "bad_user",
				os_tenant_name: "admin",
				os_password:    "secret123",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("bad tenant", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "http://10.7.0.54:5000/v2.0",
				os_username:    "admin",
				os_tenant_name: "bad_tenant",
				os_password:    "secret123",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("bad password", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "http://10.7.0.54:5000/v2.0",
				os_username:    "admin",
				os_tenant_name: "admin",
				os_password:    "letmein",
				os_token:       "",
			},
			shouldErr: true,
		}),
		Entry("bad token", TestCase{
			keys: KeystoneEnvs{
				os_auth_url:    "http://10.7.0.54:5000/v2.0",
				os_username:    "admin",
				os_tenant_name: "admin",
				os_password:    "secret123",
				os_token:       "124123412412341234",
			},
			shouldErr: true,
		}),
		Entry("everything correct", TestCase{
			// we're assuming that keystone auth from env is correct for this test.
			keys:      *TestKeystoneEnvs(),
			shouldErr: false,
		}),
	)
})
