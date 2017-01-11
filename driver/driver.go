// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package driver

import (
	"fmt"

	"github.com/Microsoft/hcsshim"
	"github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/common"
	"github.com/codilime/contrail-windows-docker/controller"
	"github.com/codilime/contrail-windows-docker/hns"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/go-plugins-helpers/sdk"
)

type ContrailDriver struct {
	controller *controller.Controller
	HnsID      string
}

func NewDriver(subnet, gateway, adapter, controllerIP string, controllerPort int) (*ContrailDriver,
	error) {

	subnets := []hcsshim.Subnet{
		{
			AddressPrefix:  subnet,
			GatewayAddress: gateway,
		},
	}

	configuration := &hcsshim.HNSNetwork{
		Name:               common.NetworkHNSname,
		Type:               "transparent",
		Subnets:            subnets,
		NetworkAdapterName: adapter,
	}

	hnsID, err := hns.CreateHNSNetwork(configuration)
	if err != nil {
		return nil, err
	}

	d := &ContrailDriver{
		controller: controller.NewController(controllerIP, controllerPort),
		HnsID:      hnsID,
	}
	return d, nil
}

func (d *ContrailDriver) Serve() error {
	h := network.NewHandler(d)

	config := sdk.WindowsPipeConfig{
		// This will set permissions for Service, System, Adminstrator group and account to have full access
		SecurityDescriptor: "D:(A;ID;FA;;;SY)(A;ID;FA;;;BA)(A;ID;FA;;;LA)(A;ID;FA;;;LS)",

		InBufferSize:  4096,
		OutBufferSize: 4096,
	}

	h.ServeWindows("//./pipe/"+common.DriverName, common.DriverName, &config)
	return nil
}

func (d *ContrailDriver) Teardown() error {
	err := hns.DeleteHNSNetwork(d.HnsID)
	return err
}

func (d *ContrailDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	logrus.Debugln("=== GetCapabilities")
	r := &network.CapabilitiesResponse{}
	r.Scope = network.LocalScope
	return r, nil
}

func (d *ContrailDriver) CreateNetwork(req *network.CreateNetworkRequest) error {
	logrus.Debugln("=== CreateNetwork")
	logrus.Debugln("network.NetworkID =", req.NetworkID)
	logrus.Debugln(req)
	logrus.Debugln("IPv4:")
	for _, n := range req.IPv4Data {
		logrus.Debugln(n)
	}
	logrus.Debugln("IPv6:")
	for _, n := range req.IPv6Data {
		logrus.Debugln(n)
	}
	logrus.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}

	return nil
}

func (d *ContrailDriver) AllocateNetwork(req *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	logrus.Debugln("=== AllocateNetwork")
	logrus.Debugln(req)
	r := &network.AllocateNetworkResponse{}
	return r, nil
}

func (d *ContrailDriver) DeleteNetwork(req *network.DeleteNetworkRequest) error {
	logrus.Debugln("=== DeleteNetwork")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) FreeNetwork(req *network.FreeNetworkRequest) error {
	logrus.Debugln("=== FreeNetwork")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) CreateEndpoint(req *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	logrus.Debugln("=== CreateEndpoint")
	logrus.Debugln(req)
	logrus.Debugln(req.Interface)
	logrus.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}
	r := &network.CreateEndpointResponse{}
	return r, nil
}

func (d *ContrailDriver) DeleteEndpoint(req *network.DeleteEndpointRequest) error {
	logrus.Debugln("=== DeleteEndpoint")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) EndpointInfo(req *network.InfoRequest) (*network.InfoResponse, error) {
	logrus.Debugln("=== EndpointInfo")
	logrus.Debugln(req)
	r := &network.InfoResponse{}
	return r, nil
}

func (d *ContrailDriver) Join(req *network.JoinRequest) (*network.JoinResponse, error) {
	logrus.Debugln("=== Join")
	logrus.Debugln(req)
	logrus.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}
	r := &network.JoinResponse{}
	return r, nil
}

func (d *ContrailDriver) Leave(req *network.LeaveRequest) error {
	logrus.Debugln("=== Leave")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) DiscoverNew(req *network.DiscoveryNotification) error {
	logrus.Debugln("=== DiscoverNew")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) DiscoverDelete(req *network.DiscoveryNotification) error {
	logrus.Debugln("=== DiscoverDelete")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) ProgramExternalConnectivity(req *network.ProgramExternalConnectivityRequest) error {
	logrus.Debugln("=== ProgramExternalConnectivity")
	logrus.Debugln(req)
	return nil
}

func (d *ContrailDriver) RevokeExternalConnectivity(req *network.RevokeExternalConnectivityRequest) error {
	logrus.Debugln("=== RevokeExternalConnectivity")
	logrus.Debugln(req)
	return nil
}
