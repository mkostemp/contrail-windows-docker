package controller

import (
	"errors"
	"fmt"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	log "github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/common"
)

type Info struct {
}

type Controller struct {
	ApiClient contrail.ApiClient
}

func NewController(ip string, port int) (*Controller, error) {
	client := &Controller{}
	client.ApiClient = contrail.NewClient(ip, port)

	// TODO: use environment variables for keystone auth (JW-66)
	keystone := contrail.NewKeystoneClient("http://10.7.0.54:5000/v2.0", "admin", "admin",
		"secret123", "")
	err := keystone.Authenticate()
	if err != nil {
		log.Errorln("Keystone error:", err)
		return nil, err
	}
	client.ApiClient.(*contrail.Client).SetAuthenticator(keystone)
	return client, nil
}

func (c *Controller) GetNetwork(tenantName, networkName string) (*types.VirtualNetwork,
	error) {
	name := fmt.Sprintf("%s:%s:%s", common.DomainName, tenantName, networkName)
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
	instance, err := types.VirtualMachineByName(c.ApiClient, containerId)
	if err == nil && instance != nil {
		return instance, nil
	}

	instance = new(types.VirtualMachine)
	instance.SetName(containerId)
	err = c.ApiClient.Create(instance)
	if err != nil {
		return nil, err
	}

	createdInstance, err := types.VirtualMachineByName(c.ApiClient, containerId)
	if err != nil {
		return nil, err
	}
	return createdInstance, nil
}

func (c *Controller) GetOrCreateInterface(net *types.VirtualNetwork,
	instance *types.VirtualMachine) (*types.VirtualMachineInterface, error) {
	iface, err := types.VirtualMachineInterfaceByName(c.ApiClient, instance.GetName())
	if err == nil && iface != nil {
		return iface, nil
	}

	iface = new(types.VirtualMachineInterface)
	instanceFQName := instance.GetFQName()
	iface.SetFQName("", instanceFQName)
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

	createdIface, err := types.VirtualMachineInterfaceByName(c.ApiClient, instance.GetName())
	if err != nil {
		return nil, err
	}
	return createdIface, nil
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
	instIp, err := types.InstanceIpByName(c.ApiClient, iface.GetName())
	if err == nil && instIp != nil {
		return instIp, nil
	}

	instIp = &types.InstanceIp{}
	instIp.SetName(iface.GetName())
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

	allocatedIP, err := types.InstanceIpByUuid(c.ApiClient, instIp.GetUuid())
	if err != nil {
		return nil, err
	}
	return allocatedIP, nil
}
