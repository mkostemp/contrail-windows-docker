package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/controller"
	"github.com/codilime/contrail-windows-docker/driver"
)

func main() {
	var adapter = flag.String("adapter", "Ethernet0",
		"net adapter for HNS switch, must be physical")
	var controllerIP = flag.String("controllerIP", "127.0.0.1",
		"IP address of Contrail Controller API")
	var controllerPort = flag.Int("controllerPort", 8082,
		"port of Contrail Controller API")
	flag.Parse()

	var d *driver.ContrailDriver
	var c *controller.Controller
	var err error

	if c, err = controller.NewController(*controllerIP, *controllerPort); err != nil {
		log.Error(err)
		return
	}

	if d, err = driver.NewDriver(*adapter, c); err != nil {
		log.Error(err)
		return
	}

	if err = d.StartServing(); err != nil {
		log.Error(err)
	}
	defer d.StopServing()
}
