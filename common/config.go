package common

import (
	"os"
	"path/filepath"
)

const (
	// DomainName specifies domain name in Contrail
	DomainName = "default-domain"

	// DriverName is name of the driver that is to be specified during docker network creation
	DriverName = "Contrail"

	// HNSNetworkPrefix is a prefix given too all HNS network names managed by the driver
	HNSNetworkPrefix = "Contrail"
)

// PluginSpecDir returns path to directory where docker daemon looks for plugin spec files.
func PluginSpecDir() string {
	return ([]string{filepath.Join(os.Getenv("programdata"), "docker", "plugins")})[0]
}

// PluginSpecFilePath returns path to plugin spec file.
func PluginSpecFilePath() string {
	return filepath.Join(PluginSpecDir(), DriverName+".spec")
}
