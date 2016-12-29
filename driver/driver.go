// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package driver

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/network"
)

const (
	// DriverName is name of the driver that is to be specified during docker network creation
	DriverName = "Contrail"

	// HNSNetworkName is a constant name for HNS network
	NetworkHNSname = "ContrailHNSNet"
)

type ContrailDriver struct {
}

func NewDriver() (*ContrailDriver, error) {
	d := &ContrailDriver{}
	return d, nil
}

func (d *ContrailDriver) Teardown() error {
	return nil
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
	r.DisableGatewayService = true
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
