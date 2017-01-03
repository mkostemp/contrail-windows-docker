package driver

import (
	"errors"
	"fmt"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
)

const (
	DomainName = "default-domain"
)

type Info struct {
}

type Controller struct {
	ApiClient contrail.ApiClient
}

func NewController(ip string, port int) *Controller {
	client := &Controller{}
	client.ApiClient = contrail.NewClient(ip, port)
	return client
}

func (c *Controller) GetNetwork(tenantName, networkName string) (*types.VirtualNetwork,
	error) {
	name := fmt.Sprintf("%s:%s:%s", DomainName, tenantName, networkName)
	net, err := types.VirtualNetworkByName(c.ApiClient, name)
	if err != nil {
		return nil, err
	}
	return net, nil
}

func (c *Controller) GetDefaultGatewayIp(net *types.VirtualNetwork) (string, error) {
	ipamReferences, err := net.GetNetworkIpamRefs()
	if err != nil {
		return "", err
	}
	if len(ipamReferences) == 0 {
		return "", errors.New("Ipam references list is empty")
	}
	attribute := ipamReferences[0].Attr
	ipamSubnets := attribute.(types.VnSubnetsType).IpamSubnets
	if len(ipamSubnets) == 0 {
		return "", errors.New("Ipam subnets list is empty")
	}
	gw := ipamSubnets[0].DefaultGateway
	if gw == "" {
		return "", errors.New("Default GW is empty")
	}
	return gw, nil
}

func (c *Controller) GetOrCreateInstance(tenantName, containerId string) (*types.VirtualMachine, error) {
	name := fmt.Sprintf("%s:%s:%s", DomainName, tenantName, containerId)
	instance, err := types.VirtualMachineByName(c.ApiClient, name)
	if err == nil && instance != nil {
		return instance, nil
	}

	instance = new(types.VirtualMachine)
	instance.SetFQName("project", []string{DomainName, tenantName, containerId})
	err = c.ApiClient.Create(instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (c *Controller) GetOrCreateInterface(net *types.VirtualNetwork,
	instance *types.VirtualMachine) (*types.VirtualMachineInterface, error) {
	instanceFQName := instance.GetFQName()
	namespace := instanceFQName[len(instanceFQName)-2]
	name := fmt.Sprintf("%s:%s:%s", DomainName, namespace, instance.GetName())
	iface, err := types.VirtualMachineInterfaceByName(c.ApiClient, name)
	if err == nil && iface != nil {
		return iface, nil
	}

	iface = new(types.VirtualMachineInterface)
	iface.SetFQName("project", []string{DomainName, namespace, instance.GetName()})
	err = iface.AddVirtualMachine(instance)
	if err != nil {
		return nil, err
	}
	err = iface.AddVirtualNetwork(net)
	if err != nil {
		return nil, err
	}
	err = c.ApiClient.Create(iface)
	if err != nil {
		return nil, err
	}
	return iface, nil
}

func (c *Controller) GetInterfaceMac(iface *types.VirtualMachineInterface) (string, error) {
	macs := iface.GetVirtualMachineInterfaceMacAddresses()
	if len(macs.MacAddress) == 0 {
		return "", errors.New("Empty MAC list")
	}
	return macs.MacAddress[0], nil
}

func (c *Controller) GetOrCreateInstanceIp(net *types.VirtualNetwork,
	iface *types.VirtualMachineInterface) (*types.InstanceIp, error) {
	ifaceFQName := iface.GetFQName()
	tenantName := ifaceFQName[len(ifaceFQName)-2]
	name := fmt.Sprintf("%s_%s", tenantName, iface.GetName())
	instIp, err := types.InstanceIpByName(c.ApiClient, name)
	if err == nil && instIp != nil {
		return instIp, nil
	}

	instIp = &types.InstanceIp{}
	instIp.SetName(name)
	err = instIp.AddVirtualNetwork(net)
	if err != nil {
		return nil, err
	}
	err = instIp.AddVirtualMachineInterface(iface)
	if err != nil {
		return nil, err
	}
	err = c.ApiClient.Create(instIp)
	if err != nil {
		return nil, err
	}
	return instIp, nil
}
