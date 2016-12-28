package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/driver"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/go-plugins-helpers/sdk"
)

func main() {
	var subnet = flag.String("subnet", "172.117.0.0/16", "subnet in CIDR format for HNS")
	var gateway = flag.String("gateway", "172.117.0.1", "default gateway IP for HNS")
	var adapter = flag.String("adapter", "Ethernet0",
		"net adapter for HNS switch, must be physical")
	flag.Parse()

	var d *driver.ContrailDriver
	var err error

	if d, err = driver.NewDriver(*subnet, *gateway, *adapter); err != nil {
		logrus.Error(err)
	}
	h := network.NewHandler(d)

	config := sdk.WindowsPipeConfig{
		// This will set permissions for Service, System, Adminstrator group and account to have full access
		SecurityDescriptor: "D:(A;ID;FA;;;SY)(A;ID;FA;;;BA)(A;ID;FA;;;LA)(A;ID;FA;;;LS)",

		InBufferSize:  4096,
		OutBufferSize: 4096,
	}

	h.ServeWindows("//./pipe/"+driver.DriverName, driver.DriverName, &config)
	err = d.Teardown()
	if err != nil {
		logrus.Error(err)
	}
}
