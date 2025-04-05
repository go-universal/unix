package systemd

import (
	"fmt"
	"os/exec"
)

const serviceTemplate = `[Unit]
Description={name}
ConditionPathExists={root}
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024

Restart=on-failure
RestartSec=10

WorkingDirectory={root}
ExecStart=/usr/bin/sudo {root}/{command}

PermissionsStartOnly=true
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier={name}

[Install]
WantedBy=multi-user.target`

// cmdError handles execution errors, including extracting the exit code and stderr output.
func cmdError(err error) error {
	if err == nil {
		return nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return fmt.Errorf("exit %d, %s", exitErr.ExitCode(), string(exitErr.Stderr))
	}

	return err
}

// reload reloads the services.
func reload() error {
	return cmdError(exec.Command("sudo", "systemctl", "daemon-reload").Run())
}
