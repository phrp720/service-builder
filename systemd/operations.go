package systemd

import (
	"fmt"
	"github.com/phrp720/service-builder/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CreateServiceFile generates the content of the .service file
func CreateServiceFile(units map[string]interface{}, service map[string]interface{}, install map[string]interface{}) string {
	var sb strings.Builder
	if len(units) > 0 {
		sb.WriteString("[Unit]\n")
		for key, value := range units {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
		sb.WriteString("\n")

	}
	if len(service) > 0 {
		sb.WriteString("[Service]\n")
		for key, value := range service {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
		sb.WriteString("\n")
	}
	if len(install) > 0 {
		sb.WriteString("[Install]\n")
		for key, value := range install {
			sb.WriteString(fmt.Sprintf("%s=%s\n", key, value))

		}
	}

	return sb.String()
}

// ToMaps converts the ServiceConfig struct to a map of unit, service and install
func (s ServiceConfig) ToMaps() (map[string]interface{}, map[string]interface{}, map[string]interface{}) {
	unitMap := util.ToMap(s.Unit)
	serviceMap := util.ToMap(s.Service)
	installMap := util.ToMap(s.Install)
	return unitMap, serviceMap, installMap
}

// CreateService writes the .service file content to the specified path
func CreateService(s ServiceConfig, path string) error {
	content := CreateServiceFile(s.ToMaps())
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0644)
}

// CreateServiceAndStart writes the .service file content to /etc/systemd/system/ and starts the service
func CreateServiceAndStart(s ServiceConfig, path string, enable bool, superUser bool) error {
	if err := CreateService(s, path); err != nil {
		return err
	}

	if enable {
		// Enable the service
		if err := EnableService(path, superUser); err != nil {
			return err
		}
	}

	// Start the service
	if err := StartService(path, superUser); err != nil {
		return err
	}
	return nil
}

// StartService enables the service
func StartService(file string, superUser bool) error {
	if err := systemdCommand("start", file, superUser); err != nil {
		return err
	}
	return nil
}

// StopService stops the service
func StopService(file string, superUser bool) error {
	if err := systemdCommand("stop", file, superUser); err != nil {
		return err
	}
	return nil
}

// EnableService enables the service
func EnableService(file string, superUser bool) error {
	if err := systemdCommand("enable", file, superUser); err != nil {
		return err
	}
	return nil
}

// DisableService disables the service
func DisableService(file string, superUser bool) error {
	if err := systemdCommand("disable", file, superUser); err != nil {
		return err
	}
	return nil
}

// RestartService restarts the service
func RestartService(file string, superUser bool) error {
	if err := systemdCommand("restart", file, superUser); err != nil {
		return err
	}
	return nil
}

// systemdCommand executes the systemctl command.If isSuperUser is true, it will run the command as sudo else as the current user
func systemdCommand(command string, file string, isSuperUser bool) error {
	var cmd *exec.Cmd
	if isSuperUser {
		cmd = exec.Command("sudo", "systemctl", command, file)
	} else {
		cmd = exec.Command("systemctl --user", command, file)
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
