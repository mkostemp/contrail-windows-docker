package common

import (
	"errors"
	"fmt"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func HardResetHNS() error {
	log.Infoln("Resetting HNS")
	log.Debugln("Removing NAT")
	if err := exec.Command("powershell", "Get-NetNat", "|",
		"Remove-NetNat").Run(); err != nil {
		log.Debugln("Could not remove container network.")
	}
	log.Debugln("Removing container networks")
	if err := exec.Command("powershell", "Get-ContainerNetwork", "|",
		"Remove-ContainerNetwork", "-Force").Run(); err != nil {
		log.Debugln("Could not remove container network.")
	}
	log.Debugln("Stopping HNS")
	if err := exec.Command("powershell", "Stop-Service", "hns").Run(); err != nil {
		log.Debugln("HNS is already stopped.")
	}
	log.Debugln("Removing HNS program data")
	if err := exec.Command("powershell", "Remove-Item",
		"C:\\ProgramData\\Microsoft\\Windows\\HNS\\HNS.data").Run(); err != nil {
		return errors.New(fmt.Sprintf("Error during removing HNS program data: %s", err))
	}
	log.Debugln("Starting HNS")
	if err := exec.Command("powershell", "Start-Service", "hns").Run(); err != nil {
		return errors.New(fmt.Sprintf("Error when starting HNS: %s", err))
	}
	return nil
}

func RestartDocker() error {
	log.Infoln("Restarting docker")
	if err := exec.Command("powershell", "Restart-Service", "docker").Run(); err != nil {
		return err
	}
	return nil
}
