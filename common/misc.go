package common

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

func HardResetHNS() error {
	log.Debugln("Removing container networks")
	if err := exec.Command("powershell", "Get-ContainerNetwork", "|",
		"Remove-ContainerNetwork").Run(); err != nil {
		log.Debugln("Could not remove container network.")
	}
	log.Debugln("Stopping HNS")
	if err := exec.Command("powershell", "Stop-Service", "hns").Run(); err != nil {
		log.Debugln("HNS is already stopped.")
	}
	log.Debugln("Removing HNS program data")
	if err := exec.Command("powershell", "Remove-Item",
		"C:\\ProgramData\\Microsoft\\Windows\\HNS\\HNS.data").Run(); err != nil {
		return err
	}
	log.Debugln("Starting HNS")
	if err := exec.Command("powershell", "Start-Service", "hns").Run(); err != nil {
		return err
	}
	return nil
}
