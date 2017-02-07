package common

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
		log.Debugln("Could not remove nat network.")
	}
	log.Debugln("Stopping HNS")
	if err := exec.Command("powershell", "Stop-Service", "hns").Run(); err != nil {
		log.Debugln("HNS is already stopped.")
	}
	log.Debugln("Removing HNS program data")

	programData := os.Getenv("programdata")
	if programData == "" {
		log.Errorln("Invalid program data env variable")
		return errors.New("Invalid program data env variable")
	}
	hnsDataDir := filepath.Join(programData, "Microsoft", "Windows", "HNS", "HNS.data")
	if err := exec.Command("powershell", "Remove-Item", hnsDataDir).Run(); err != nil {
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
