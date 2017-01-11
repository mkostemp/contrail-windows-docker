package controller

import (
	"testing"

	"github.com/Juniper/contrail-go-api/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

var _ = Describe("Controller", func() {

	var client *Controller
	var project *types.Project

	BeforeEach(func() {
		client, project = NewMockedClientAndProject(tenantName)
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
				Expect(net).To(BeEquivalentTo(testNetwork))
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

	Describe("getting Contrail default gateway IP", func() {
		Context("network has subnet with default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
				AddSubnetWithDefaultGateway(client.ApiClient, subnetPrefix, defaultGW,
					subnetMask, testNetwork)
			})
			It("returns gateway IP", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				Expect(err).ToNot(HaveOccurred())
				Expect(gwAddr).ToNot(Equal(""))
			})
		})
		Context("network has subnet without default gateway", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetworkWithSubnet(client.ApiClient, networkName,
					subnetCIDR, project)
			})
			It("returns error", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				Expect(err).To(HaveOccurred())
				Expect(gwAddr).To(Equal(""))
			})
		})
		Context("network doesn't have subnets", func() {
			var testNetwork *types.VirtualNetwork
			BeforeEach(func() {
				testNetwork = CreateMockedNetwork(client.ApiClient, networkName, project)
			})
			It("returns error", func() {
				gwAddr, err := client.GetDefaultGatewayIp(testNetwork)
				Expect(err).To(HaveOccurred())
				Expect(gwAddr).To(Equal(""))
			})
		})
	})

	Describe("getting Contrail instance", func() {
		Context("when instance already exists in Cotnrail", func() {
			var testInstance *types.VirtualMachine
			BeforeEach(func() {
				testInstance = CreateMockedInstance(client.ApiClient, tenantName, containerID)
			})
			It("returns existing instance", func() {
				instance, err := client.GetOrCreateInstance(tenantName, containerID)
				Expect(err).ToNot(HaveOccurred())
				Expect(instance).ToNot(BeNil())
				Expect(instance).To(BeEquivalentTo(testInstance))
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
				Expect(existingInst).To(BeEquivalentTo(instance))
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
				Expect(iface).To(BeEquivalentTo(testInterface))
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
				Expect(existingIface).To(BeEquivalentTo(iface))
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
				Expect(instanceIP).To(BeEquivalentTo(testInstanceIP))

				// TODO: check if got an IP address

			})
		})
		Context("when instance IP doesn't exist in Contrail", func() {
			It("creates new instance IP", func() {
				instanceIP, err := client.GetOrCreateInstanceIp(testNetwork, testInterface)
				Expect(err).ToNot(HaveOccurred())
				Expect(instanceIP).ToNot(BeNil())

				existingIP, err := types.InstanceIpByUuid(client.ApiClient, instanceIP.GetUuid())
				Expect(err).ToNot(HaveOccurred())
				Expect(existingIP).To(BeEquivalentTo(instanceIP))

				// TODO: check if got an IP address
			})
		})
	})
})
