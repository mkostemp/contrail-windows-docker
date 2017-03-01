package controller

import (
	"flag"
	"testing"

	contrail "github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	log "github.com/Sirupsen/logrus"

	"github.com/codilime/contrail-windows-docker/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/onsi/ginkgo/reporters"
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
	// asdf

	log.SetLevel(log.DebugLevel)
}

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("controller_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Controller client test suite",
		[]Reporter{junitReporter})
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
		testInterface := CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
			containerID)
		_ = CreateMockedInstance(client.ApiClient, testInterface, containerID)
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
		testInterface := CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
			containerID)
		testInstance := CreateMockedInstance(client.ApiClient, testInterface, containerID)
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
		Context("network has subnet with default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
				AddSubnetWithDefaultGateway(client.ApiClient, subnetPrefix, defaultGW,
					subnetMask, testNetwork)
			})
			Specify("getting default gw IP works", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				Expect(err).ToNot(HaveOccurred())
				Expect(gwAddr).ToNot(Equal(""))
			})
			Specify("getting subnet prefix and prefix len works", func() {
				ipam, err := client.GetIpamSubnet(testNetwork)
				Expect(err).ToNot(HaveOccurred())
				Expect(ipam.Subnet.IpPrefix).To(Equal(subnetPrefix))
				Expect(ipam.Subnet.IpPrefixLen).To(Equal(subnetMask))
			})
		})
		Context("network has subnet without default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName,
					subnetCIDR, project)
			})
			Specify("getting default gw IP returns error", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				if useActualController {
					Expect(gwAddr).ToNot(Equal(""))
					Expect(err).ToNot(HaveOccurred())
				} else {
					// mocked controller lacks some logic here
					Expect(gwAddr).To(Equal(""))
					Expect(err).To(HaveOccurred())
				}
			})
			Specify("getting subnet prefix and prefix len works", func() {
				ipam, err := client.GetIpamSubnet(testNetwork)
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
			Specify("getting default gw IP returns error", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				Expect(err).To(HaveOccurred())
				Expect(gwAddr).To(Equal(""))
			})
			Specify("getting subnet prefix and prefix len returns error", func() {
				ipam, err := client.GetIpamSubnet(testNetwork)
				Expect(err).To(HaveOccurred())
				Expect(ipam).To(BeNil())
			})
		})
	})

	Describe("getting Contrail virtual interface", func() {
		var testNetwork *types.VirtualNetwork
		BeforeEach(func() {
			testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
		})
		Context("when vif already exists in Contrail", func() {
			var testInterface *types.VirtualMachineInterface
			BeforeEach(func() {
				testInterface = CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
					containerID)
			})
			It("returns existing vif", func() {
				iface, err := client.GetOrCreateInterface(testNetwork, tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(iface).ToNot(BeNil())
				Expect(iface.GetUuid()).To(Equal(testInterface.GetUuid()))
			})
			It("assigns correct FQName to vif", func() {
				iface, err := client.GetOrCreateInterface(testNetwork, tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(iface).ToNot(BeNil())
				Expect(iface.GetFQName()).To(Equal([]string{common.DomainName, tenantName,
					containerID}))
			})
		})
		Context("when vif doesn't exist in Contrail", func() {
			It("creates a new vif", func() {
				iface, err := client.GetOrCreateInterface(testNetwork, tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(iface).ToNot(BeNil())

				existingIface, err := types.VirtualMachineInterfaceByUuid(client.ApiClient,
					iface.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingIface.GetUuid()).To(Equal(iface.GetUuid()))
			})
		})
	})

	Describe("getting Contrail instance", func() {
		var testInterface *types.VirtualMachineInterface
		BeforeEach(func() {
			testNetwork := CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
			testInterface = CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
				containerID)
		})
		Context("when instance already exists in Contrail", func() {
			var testInstance *types.VirtualMachine
			BeforeEach(func() {
				testInstance = CreateMockedInstance(client.ApiClient, testInterface, containerID)
			})
			It("returns existing instance", func() {
				instance, err := client.GetOrCreateInstance(testInterface, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(instance).ToNot(BeNil())
				Expect(instance.GetUuid()).To(Equal(testInstance.GetUuid()))
			})
		})
		Context("when instance doesn't exist in Contrail", func() {
			It("creates a new instance", func() {
				instance, err := client.GetOrCreateInstance(testInterface, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(instance).ToNot(BeNil())

				existingInst, err := types.VirtualMachineByUuid(client.ApiClient,
					instance.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingInst.GetUuid()).To(Equal(instance.GetUuid()))
			})
		})
	})

	Describe("getting virtual interface MAC", func() {
		var testInterface *types.VirtualMachineInterface
		BeforeEach(func() {
			testNetwork := CreateMockedNetworkWithSubnet(client.ApiClient, networkName, subnetCIDR,
				project)
			testInterface = CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
				containerID)
		})
		Context("when vif has a VM", func() {
			BeforeEach(func() {
				_ = CreateMockedInstance(client.ApiClient, testInterface, containerID)
			})
			It("returns MAC address", func() {
				mac, err := client.GetInterfaceMac(testInterface)
				Expect(err).ToNot(HaveOccurred())
				Expect(mac).ToNot(Equal("")) // dunno how to get actual MAC when given Instance
			})
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
			testInterface = CreateMockedInterface(client.ApiClient, testNetwork, tenantName,
				containerID)
			testInstance = CreateMockedInstance(client.ApiClient, testInterface, containerID)
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
