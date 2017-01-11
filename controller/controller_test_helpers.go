package controller

import (
	"fmt"

	contrail "github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/config"
	"github.com/Juniper/contrail-go-api/mocks"
	"github.com/Juniper/contrail-go-api/types"
	"github.com/codilime/contrail-windows-docker/common"
	. "github.com/onsi/gomega"
)

func NewMockedClientAndProject(tenant string) (*Controller, *types.Project) {
	c := &Controller{}
	mockedApiClient := new(mocks.ApiClient)
	mockedApiClient.Init()
	c.ApiClient = mockedApiClient

	project := new(types.Project)
	project.SetFQName("domain", []string{common.DomainName, tenant})
	err := c.ApiClient.Create(project)
	Expect(err).ToNot(HaveOccurred())
	return c, project
}

func CreateMockedNetworkWithSubnet(c contrail.ApiClient, netName, subnetCIDR string,
	project *types.Project) *types.VirtualNetwork {
	netUUID, err := config.CreateNetworkWithSubnet(c, project.GetUuid(), netName, subnetCIDR)
	Expect(err).ToNot(HaveOccurred())
	Expect(netUUID).ToNot(BeNil())
	testNetwork, err := types.VirtualNetworkByUuid(c, netUUID)
	Expect(err).ToNot(HaveOccurred())
	Expect(testNetwork).ToNot(BeNil())
	return testNetwork
}

func CreateMockedNetwork(c contrail.ApiClient, netName string,
	project *types.Project) *types.VirtualNetwork {
	netUUID, err := config.CreateNetwork(c, project.GetUuid(), netName)
	Expect(err).ToNot(HaveOccurred())
	Expect(netUUID).ToNot(BeNil())
	testNetwork, err := types.VirtualNetworkByUuid(c, netUUID)
	Expect(err).ToNot(HaveOccurred())
	Expect(testNetwork).ToNot(BeNil())
	return testNetwork
}

func AddSubnetWithDefaultGateway(c contrail.ApiClient, subnetPrefix, defaultGW string,
	subnetMask int, testNetwork *types.VirtualNetwork) {
	subnet := &types.IpamSubnetType{
		Subnet:         &types.SubnetType{IpPrefix: subnetPrefix, IpPrefixLen: subnetMask},
		DefaultGateway: defaultGW,
	}

	var ipamSubnets types.VnSubnetsType
	ipamSubnets.AddIpamSubnets(subnet)

	ipam, err := c.FindByName("network-ipam",
		"default-domain:default-project:default-network-ipam")
	Expect(err).ToNot(HaveOccurred())
	err = testNetwork.AddNetworkIpam(ipam.(*types.NetworkIpam), ipamSubnets)
	Expect(err).ToNot(HaveOccurred())

	err = c.Update(testNetwork)
	Expect(err).ToNot(HaveOccurred())
}

func CreateMockedInstance(c contrail.ApiClient, tenantName,
	containerID string) *types.VirtualMachine {
	testInstance := new(types.VirtualMachine)
	testInstance.SetFQName("project", []string{common.DomainName, tenantName, containerID})
	err := c.Create(testInstance)
	Expect(err).ToNot(HaveOccurred())
	return testInstance
}

func CreateMockedInterface(c contrail.ApiClient, instance *types.VirtualMachine,
	net *types.VirtualNetwork) *types.VirtualMachineInterface {
	iface := new(types.VirtualMachineInterface)
	instanceFQName := instance.GetFQName()
	namespace := instanceFQName[len(instanceFQName)-2]
	iface.SetFQName("project", []string{common.DomainName, namespace, instance.GetName()})
	err := iface.AddVirtualMachine(instance)
	Expect(err).ToNot(HaveOccurred())
	err = iface.AddVirtualNetwork(net)
	Expect(err).ToNot(HaveOccurred())
	err = c.Create(iface)
	Expect(err).ToNot(HaveOccurred())
	return iface
}

func AddMacToInterface(c contrail.ApiClient, ifaceMac string,
	iface *types.VirtualMachineInterface) {
	macs := new(types.MacAddressesType)
	macs.AddMacAddress(ifaceMac)
	iface.SetVirtualMachineInterfaceMacAddresses(macs)
	err := c.Update(iface)
	Expect(err).ToNot(HaveOccurred())
}

func CreateMockedInstanceIP(c contrail.ApiClient, tenantName string,
	iface *types.VirtualMachineInterface,
	net *types.VirtualNetwork) *types.InstanceIp {
	name := fmt.Sprintf("%s_%s", tenantName, iface.GetName())

	instIP := &types.InstanceIp{}
	instIP.SetName(name)
	err := instIP.AddVirtualNetwork(net)
	Expect(err).ToNot(HaveOccurred())
	err = instIP.AddVirtualMachineInterface(iface)
	Expect(err).ToNot(HaveOccurred())
	err = c.Create(instIP)
	Expect(err).ToNot(HaveOccurred())

	allocatedIP, err := types.InstanceIpByUuid(c, instIP.GetUuid())
	Expect(err).ToNot(HaveOccurred())
	return allocatedIP
}
