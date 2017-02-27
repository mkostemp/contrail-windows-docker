package common

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

func callPowershell(cmds ...string) error {
	c := []string{"-NonInteractive"}
	for _, cmd := range cmds {
		c = append(c, cmd)
	}
	return exec.Command("powershell", c...).Run()
}

func HardResetHNS() error {
	log.Infoln("Resetting HNS")
	log.Debugln("Removing NAT")
	if err := callPowershell("Get-NetNat", "|", "Remove-NetNat"); err != nil {
		log.Debugln("Could not remove nat network.")
	}
	log.Debugln("Removing container networks")
	if err := callPowershell("Get-ContainerNetwork", "|", "Remove-ContainerNetwork",
		"-Force"); err != nil {
		log.Debugln("Could not remove container network.")
	}
	log.Debugln("Stopping HNS")
	if err := callPowershell("Stop-Service", "hns"); err != nil {
		log.Debugln("HNS is already stopped.")
	}
	log.Debugln("Removing HNS program data")

	programData := os.Getenv("programdata")
	if programData == "" {
		return errors.New("Invalid program data env variable")
	}
	hnsDataDir := filepath.Join(programData, "Microsoft", "Windows", "HNS", "HNS.data")
	if err := callPowershell("Remove-Item", hnsDataDir); err != nil {
		return errors.New(fmt.Sprintf("Error during removing HNS program data: %s", err))
	}
	log.Debugln("Starting HNS")
	if err := callPowershell("Start-Service", "hns"); err != nil {
		return errors.New(fmt.Sprintf("Error when starting HNS: %s", err))
	}
	return nil
}

func RestartDocker() error {
	log.Infoln("Restarting docker")
	if err := callPowershell("Restart-Service", "docker"); err != nil {
		return err
	}
	return nil
}
