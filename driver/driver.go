// Implemented according to
// https://github.com/docker/libnetwork/blob/master/docs/remote.md

package driver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"context"

	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/hcsshim"
	log "github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/common"
	"github.com/codilime/contrail-windows-docker/controller"
	"github.com/codilime/contrail-windows-docker/hns"
	"github.com/codilime/contrail-windows-docker/hnsManager"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/go-plugins-helpers/network"
)

type ContrailDriver struct {
	controller     *controller.Controller
	hnsMgr         *hnsManager.HNSManager
	networkAdapter string
	listener       net.Listener
}

func NewDriver(adapter string, c *controller.Controller) (*ContrailDriver, error) {

	d := &ContrailDriver{
		controller:     c,
		hnsMgr:         &hnsManager.HNSManager{},
		networkAdapter: adapter,
	}
	return d, nil
}

func (d *ContrailDriver) StartServing() error {

	pipeConfig := winio.PipeConfig{
		// This will set permissions for Service, System, Adminstrator group and account to
		// have full access
		SecurityDescriptor: "D:(A;ID;FA;;;SY)(A;ID;FA;;;BA)(A;ID;FA;;;LA)(A;ID;FA;;;LS)",
		MessageMode:        true,
		InputBufferSize:    4096,
		OutputBufferSize:   4096,
	}

	var err error
	pipeAddr := "//./pipe/" + common.DriverName
	if d.listener, err = winio.ListenPipe(pipeAddr, &pipeConfig); err != nil {
		return err
	}

	if err := os.MkdirAll(common.PluginSpecDir(), 0755); err != nil {
		return err
	}

	url := "npipe://" + d.listener.Addr().String()
	if err := ioutil.WriteFile(common.PluginSpecFilePath(), []byte(url), 0644); err != nil {
		return err
	}

	h := network.NewHandler(d)
	go h.Serve(d.listener)

	// wait for listener goroutine to spin up. I don't see more elegant way to do this.
	time.Sleep(time.Second * 1)

	log.Infoln("Started serving on ", pipeAddr)

	return nil
}

func (d *ContrailDriver) StopServing() error {
	_ = os.Remove(common.PluginSpecFilePath())

	if err := d.listener.Close(); err != nil {
		log.Errorln(err)
		return err
	}

	log.Infoln("Stopped serving")

	return nil
}

func (d *ContrailDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	log.Debugln("=== GetCapabilities")
	r := &network.CapabilitiesResponse{}
	r.Scope = network.LocalScope
	return r, nil
}

func (d *ContrailDriver) CreateNetwork(req *network.CreateNetworkRequest) error {
	log.Debugln("=== CreateNetwork")
	log.Debugln("network.NetworkID =", req.NetworkID)
	log.Debugln(req)
	log.Debugln("IPv4:")
	for _, n := range req.IPv4Data {
		log.Debugln(n)
	}
	log.Debugln("IPv6:")
	for _, n := range req.IPv6Data {
		log.Debugln(n)
	}
	log.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}

	reqGenericOptionsMap, exists := req.Options["com.docker.network.generic"]
	if !exists {
		return errors.New("Generic options missing")
	}

	genericOptions, ok := reqGenericOptionsMap.(map[string]interface{})
	if !ok {
		return errors.New("Malformed generic options")
	}

	tenant, exists := genericOptions["tenant"]
	if !exists {
		return errors.New("Tenant not specified")
	}

	netName, exists := genericOptions["network"]
	if !exists {
		return errors.New("Network name not specified")
	}

	// Check if network is already created in Contrail.
	contrailNetwork, err := d.controller.GetNetwork(tenant.(string), netName.(string))
	if err != nil {
		return err
	}
	if contrailNetwork == nil {
		return errors.New("Retreived Contrail network is nil")
	}

	log.Infoln("Got Contrail network", contrailNetwork.GetDisplayName())

	contrailIpam, err := d.controller.GetIpamSubnet(contrailNetwork)
	if err != nil {
		return err
	}
	subnet := contrailIpam.Subnet
	subnetCIDR := fmt.Sprintf("%s/%v", subnet.IpPrefix, subnet.IpPrefixLen)

	gw, err := d.controller.GetDefaultGatewayIp(contrailNetwork)
	if err != nil {
		return err
	}

	_, err = d.hnsMgr.CreateNetwork(d.networkAdapter, tenant.(string), netName.(string),
		subnetCIDR, gw)

	return err
}

func (d *ContrailDriver) AllocateNetwork(req *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	log.Debugln("=== AllocateNetwork")
	log.Debugln(req)
	return nil, errors.New("AllocateNetwork is not implemented")
}

func (d *ContrailDriver) DeleteNetwork(req *network.DeleteNetworkRequest) error {
	log.Debugln("=== DeleteNetwork")
	log.Debugln(req)

	tenant, netName, err := d.tenantAndNetnameFromDockerNetwork(req.NetworkID)
	if err != nil {
		return err
	}
	err = d.hnsMgr.DeleteNetwork(tenant, netName)
	if err != nil {
		return err
	}
	return nil
}

func (d *ContrailDriver) FreeNetwork(req *network.FreeNetworkRequest) error {
	log.Debugln("=== FreeNetwork")
	log.Debugln(req)
	return errors.New("FreeNetwork is not implemented")
}

func (d *ContrailDriver) CreateEndpoint(req *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	log.Debugln("=== CreateEndpoint")
	log.Debugln(req)
	log.Debugln(req.Interface)
	log.Debugln(req.EndpointID)
	log.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}

	tenant, netName, err := d.tenantAndNetnameFromDockerNetwork(req.NetworkID)
	if err != nil {
		return nil, err
	}

	contrailNetwork, err := d.controller.GetNetwork(tenant, netName)
	log.Infoln("Retreived Contrail network:", contrailNetwork.GetUuid())
	if err != nil {
		return nil, err
	}

	// TODO JW-187.
	// We need to retreive Container ID here and use it instead of EndpointID as
	// argument to d.controller.GetOrCreateInstance().
	// EndpointID is equiv to interface, but in Contrail, we have a "VirtualMachine" in
	// data model.
	// A single VM can be connected to two or more overlay networks, but when we use
	// EndpointID, this won't work.
	// We need something like:
	// containerID := req.Options["vmname"]
	containerID := req.EndpointID

	contrailInstance, err := d.controller.GetOrCreateInstance(tenant, containerID)
	if err != nil {
		return nil, err
	}

	contrailVif, err := d.controller.GetOrCreateInterface(contrailNetwork, contrailInstance)
	if err != nil {
		return nil, err
	}

	contrailIP, err := d.controller.GetOrCreateInstanceIp(contrailNetwork, contrailVif)
	log.Infoln("Retreived instance IP:", contrailIP.GetInstanceIpAddress())
	if err != nil {
		return nil, err
	}

	contrailGateway, err := d.controller.GetDefaultGatewayIp(contrailNetwork)
	log.Infoln("Retreived GW address:", contrailGateway)
	if err != nil {
		return nil, err
	}

	contrailMac, err := d.controller.GetInterfaceMac(contrailVif)
	log.Infoln("Retreived MAC:", contrailMac)
	if err != nil {
		return nil, err
	}
	// contrail MACs are like 11:22:aa:bb:cc:dd
	// HNS needs MACs like 11-22-AA-BB-CC-DD
	formattedMac := strings.Replace(strings.ToUpper(contrailMac), ":", "-", -1)

	hnsNet, err := d.hnsMgr.GetNetwork(tenant, netName)
	if err != nil {
		return nil, err
	}

	hnsEndpointConfig := &hcsshim.HNSEndpoint{
		VirtualNetworkName: hnsNet.Name,
		Name:               req.EndpointID,
		IPAddress:          net.ParseIP(contrailIP.GetInstanceIpAddress()),
		MacAddress:         formattedMac,
		GatewayAddress:     contrailGateway,
	}

	// TODO: maybe store hnsEndpointID somehow? is there a reason to?
	// Maybe it will become more clear when implementing the rest of the API.
	_, err = hns.CreateHNSEndpoint(hnsEndpointConfig)
	if err != nil {
		return nil, err
	}

	// TODO JW-12: talk to vRouter here

	contrailIpam, err := d.controller.GetIpamSubnet(contrailNetwork)
	if err != nil {
		return nil, err
	}

	epAddressCIDR := fmt.Sprintf("%s/%v", contrailIP.GetInstanceIpAddress(),
		contrailIpam.Subnet.IpPrefixLen)

	r := &network.CreateEndpointResponse{
		Interface: &network.EndpointInterface{
			Address:    epAddressCIDR,
			MacAddress: contrailMac,
		},
	}
	return r, nil
}

func (d *ContrailDriver) DeleteEndpoint(req *network.DeleteEndpointRequest) error {
	log.Debugln("=== DeleteEndpoint")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) EndpointInfo(req *network.InfoRequest) (*network.InfoResponse, error) {
	log.Debugln("=== EndpointInfo")
	log.Debugln(req)
	r := &network.InfoResponse{}
	return r, nil
}

func (d *ContrailDriver) Join(req *network.JoinRequest) (*network.JoinResponse, error) {
	log.Debugln("=== Join")
	log.Debugln(req)
	log.Debugln("options:")
	for k, v := range req.Options {
		fmt.Printf("%v: %v\n", k, v)
	}

	tenant, netName, err := d.tenantAndNetnameFromDockerNetwork(req.NetworkID)
	if err != nil {
		return nil, err
	}

	contrailNetwork, err := d.controller.GetNetwork(tenant, netName)
	log.Infoln("Retreived Contrail network:", contrailNetwork.GetUuid())
	if err != nil {
		return nil, err
	}

	gw, err := d.controller.GetDefaultGatewayIp(contrailNetwork)
	if err != nil {
		return nil, err
	}

	r := &network.JoinResponse{
		DisableGatewayService: true,
		Gateway:               gw,
	}

	return r, nil
}

func (d *ContrailDriver) Leave(req *network.LeaveRequest) error {
	log.Debugln("=== Leave")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) DiscoverNew(req *network.DiscoveryNotification) error {
	log.Debugln("=== DiscoverNew")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) DiscoverDelete(req *network.DiscoveryNotification) error {
	log.Debugln("=== DiscoverDelete")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) ProgramExternalConnectivity(req *network.ProgramExternalConnectivityRequest) error {
	log.Debugln("=== ProgramExternalConnectivity")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) RevokeExternalConnectivity(req *network.RevokeExternalConnectivityRequest) error {
	log.Debugln("=== RevokeExternalConnectivity")
	log.Debugln(req)
	return nil
}

func (d *ContrailDriver) tenantAndNetnameFromDockerNetwork(dockerNetID string) (string, string, error) {
	docker, err := dockerClient.NewEnvClient()
	if err != nil {
		return "", "", err
	}

	dockerNetwork, err := docker.NetworkInspect(context.Background(), dockerNetID)
	if err != nil {
		return "", "", err
	}

	tenant, exists := dockerNetwork.Options["tenant"]
	if !exists {
		return "", "", errors.New("Retreived network has no Contrail tenant specified")
	}

	netName, exists := dockerNetwork.Options["network"]
	if !exists {
		return "", "", errors.New("Retreived network has no Contrail network name specfied")
	}

	return tenant, netName, nil
}
