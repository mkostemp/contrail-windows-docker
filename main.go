// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package main

import (
	"encoding/json"
	"fmt"

	"github.com/Microsoft/hcsshim"
	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/go-plugins-helpers/sdk"
)

type dummyNetworkDriver struct{}

func main() {
	d := dummyNetworkDriver{}
	h := network.NewHandler(d)

	config := sdk.WindowsPipeConfig{
		// This will set permissions for Everyone user allowing him to open, write, read the pipe
		SecurityDescriptor: "S:(ML;;NW;;;LW)D:(A;;0x12019f;;;WD)",
		InBufferSize:       4096,
		OutBufferSize:      4096,
	}

	h.ServeWindows("//./pipe/testpipe", "testpipe", &config)
}

func (dummyNetworkDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	fmt.Println("GetCapabilities")
	r := &network.CapabilitiesResponse{}
	r.Scope = network.LocalScope
	return r, nil
}

func (dummyNetworkDriver) CreateNetwork(network *network.CreateNetworkRequest) error {
	fmt.Println("CreateNetwork")
	fmt.Println("network.NetworkID =", network.NetworkID)

	genericOptions := network.Options["com.docker.network.generic"].(map[string]interface{})

	vSwitch := genericOptions["vswitch"].(string)

	fmt.Println("VSwitch =", vSwitch)

	networkObject := &hcsshim.HNSNetwork{
		Id:   network.NetworkID,
		Name: "random",
	}

	request, err := json.Marshal(networkObject)

	if err != nil {
		return err
	}

	logrus.Println("Request", request)

	return nil
}

func (dummyNetworkDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	fmt.Println("AllocateNetwork")
	r := &network.AllocateNetworkResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	fmt.Println("DeleteNetwork")
	return nil
}

func (dummyNetworkDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	fmt.Println("FreeNetwork")
	return nil
}

func (dummyNetworkDriver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	fmt.Println("CreateEndpoint")
	r := &network.CreateEndpointResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	fmt.Println("DeleteEndpoint")
	return nil
}

func (dummyNetworkDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	fmt.Println("EndpointInfo")
	r := &network.InfoResponse{}
	return r, nil
}

func (dummyNetworkDriver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Println("Join")
	r := &network.JoinResponse{}
	r.DisableGatewayService = true
	return r, nil
}

func (dummyNetworkDriver) Leave(*network.LeaveRequest) error {
	fmt.Println("Leave")
	return nil
}

func (dummyNetworkDriver) DiscoverNew(*network.DiscoveryNotification) error {
	fmt.Println("DiscoverNew")
	return nil
}

func (dummyNetworkDriver) DiscoverDelete(*network.DiscoveryNotification) error {
	fmt.Println("DiscoverDelete")
	return nil
}

func (dummyNetworkDriver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	fmt.Println("ProgramExternalConnectivity")
	return nil
}

func (dummyNetworkDriver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	fmt.Println("RevokeExternalConnectivity")
	return nil
}
