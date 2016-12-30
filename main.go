package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/driver"
)

func main() {
	var subnet = flag.String("subnet", "172.117.0.0/16", "subnet in CIDR format for HNS")
	var gateway = flag.String("gateway", "172.117.0.1", "default gateway IP for HNS")
	var adapter = flag.String("adapter", "Ethernet0",
		"net adapter for HNS switch, must be physical")
	var controllerIP = flag.String("controllerIP", "127.0.0.1",
		"IP address of Contrail Controller API")
	var controllerPort = flag.Int("controllerPort", 8082,
		"port of Contrail Controller API")
	flag.Parse()

	var d *driver.ContrailDriver
	var err error

	if d, err = driver.NewDriver(*subnet, *gateway, *adapter, *controllerIP,
		*controllerPort); err != nil {
		logrus.Error(err)
		return
	}

	if err = d.Serve(); err != nil {
		logrus.Error(err)
		return
	}

	if err = d.Teardown(); err != nil {
		logrus.Error(err)
	}
}
