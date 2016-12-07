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
	fmt.Println("=== GetCapabilities")
	r := &network.CapabilitiesResponse{}
	r.Scope = network.LocalScope
	return r, nil
}

func (dummyNetworkDriver) CreateNetwork(req *network.CreateNetworkRequest) error {
	fmt.Println("=== CreateNetwork")
	fmt.Println("network.NetworkID =", req.NetworkID)
	fmt.Println(req)
	fmt.Println("IPv4:")
	for _, n := range req.IPv4Data {
		fmt.Println(n)
	}
	fmt.Println("IPv6:")
	for _, n := range req.IPv6Data {
		fmt.Println(n)
	}
	fmt.Println("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}

	network := req

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

func (dummyNetworkDriver) AllocateNetwork(req *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	fmt.Println("=== AllocateNetwork")
	fmt.Println(req)
	r := &network.AllocateNetworkResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteNetwork(req *network.DeleteNetworkRequest) error {
	fmt.Println("=== DeleteNetwork")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) FreeNetwork(req *network.FreeNetworkRequest) error {
	fmt.Println("=== FreeNetwork")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) CreateEndpoint(req *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	fmt.Println("=== CreateEndpoint")
	fmt.Println(req)
	fmt.Println(req.Interface)
	fmt.Println("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}
	r := &network.CreateEndpointResponse{}
	return r, nil
}

func (dummyNetworkDriver) DeleteEndpoint(req *network.DeleteEndpointRequest) error {
	fmt.Println("=== DeleteEndpoint")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) EndpointInfo(req *network.InfoRequest) (*network.InfoResponse, error) {
	fmt.Println("=== EndpointInfo")
	fmt.Println(req)
	r := &network.InfoResponse{}
	return r, nil
}

func (dummyNetworkDriver) Join(req *network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Println("=== Join")
	fmt.Println(req)
	fmt.Println("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}
	r := &network.JoinResponse{}
	r.DisableGatewayService = true
	return r, nil
}

func (dummyNetworkDriver) Leave(req *network.LeaveRequest) error {
	fmt.Println("=== Leave")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) DiscoverNew(req *network.DiscoveryNotification) error {
	fmt.Println("=== DiscoverNew")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) DiscoverDelete(req *network.DiscoveryNotification) error {
	fmt.Println("=== DiscoverDelete")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) ProgramExternalConnectivity(req *network.ProgramExternalConnectivityRequest) error {
	fmt.Println("=== ProgramExternalConnectivity")
	fmt.Println(req)
	return nil
}

func (dummyNetworkDriver) RevokeExternalConnectivity(req *network.RevokeExternalConnectivityRequest) error {
	fmt.Println("=== RevokeExternalConnectivity")
	fmt.Println(req)
	return nil
}
