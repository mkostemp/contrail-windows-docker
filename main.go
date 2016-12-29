package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/codilime/contrail-windows-docker/driver"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/go-plugins-helpers/sdk"
)

func main() {
	d, err := driver.NewDriver()
	if err != nil {
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
