package hnsManager

import (
	"errors"
	"fmt"

	"github.com/Microsoft/hcsshim"
	"github.com/codilime/contrail-windows-docker/common"
	"github.com/codilime/contrail-windows-docker/hns"
)

// HNSManager manages HNS networks that are used by the driver.
type HNSManager struct {
	// TODO JW-154: store networks here that you know about. If not found here, look in HNS.
	// for now, just look in HNS by name.
}

func contrailHNSNetName(tenant, netName string) string {
	return fmt.Sprintf("%s:%s:%s", common.HNSNetworkPrefix, tenant, netName)
}

func (m *HNSManager) CreateNetwork(netAdapter, tenantName, networkName, subnetCIDR,
	defaultGW string) (*hcsshim.HNSNetwork, error) {

	hnsNetName := contrailHNSNetName(tenantName, networkName)

	net, err := hns.GetHNSNetworkByName(hnsNetName)
	if net != nil {
		return nil, errors.New("Such HNS network already exists")
	}

	subnets := []hcsshim.Subnet{
		{
			AddressPrefix:  subnetCIDR,
			GatewayAddress: defaultGW,
		},
	}

	configuration := &hcsshim.HNSNetwork{
		Name:               hnsNetName,
		Type:               "transparent",
		NetworkAdapterName: netAdapter,
		Subnets:            subnets,
	}

	hnsNetworkID, err := hns.CreateHNSNetwork(configuration)
	if err != nil {
		return nil, err
	}

	hnsNetwork, err := hns.GetHNSNetwork(hnsNetworkID)
	if err != nil {
		return nil, err
	}

	return hnsNetwork, nil
}

func (m *HNSManager) GetNetwork(tenantName, networkName string) (*hcsshim.HNSNetwork, error) {
	hnsNetName := contrailHNSNetName(tenantName, networkName)
	hnsNetwork, err := hns.GetHNSNetworkByName(hnsNetName)
	if err != nil {
		return nil, err
	}
	if hnsNetwork == nil {
		return nil, errors.New("Such HNS network does not exist")
	}
	return hnsNetwork, nil
}

func (m *HNSManager) DeleteNetwork(tenantName, networkName string) error {
	hnsNetwork, err := m.GetNetwork(tenantName, networkName)
	if err != nil {
		return err
	}
	endpoints, err := hns.ListHNSEndpoints()
	if err != nil {
		return err
	}

	for _, ep := range endpoints {
		if ep.VirtualNetworkName == hnsNetwork.Name {
			return errors.New("Cannot delete network with active endpoints")
		}
	}

	return hns.DeleteHNSNetwork(hnsNetwork.Id)
}
