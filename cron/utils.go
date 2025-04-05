package cron

import (
	"fmt"
	"os/exec"
	"strings"
)

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

// parseCommand extracts the command portion from a cron expression.
// It supports both predefined constants (e.g., @daily) and custom cron expressions.
func parseCommand(cronExpr string) (bool, string) {
	aliases := []string{
		"@reboot ", "@yearly ", "@annually ",
		"@monthly ", "@weekly ", "@daily ",
		"@midnight ", "@hourly ",
	}

	for _, alias := range aliases {
		if strings.HasPrefix(cronExpr, alias) {
			return true, strings.TrimSpace(strings.TrimPrefix(cronExpr, alias))
		}
	}

	parts := strings.Fields(cronExpr)
	if len(parts) < 6 {
		return false, ""
	}

	return true, strings.TrimSpace(strings.Join(parts[5:], " "))
}

// allCrons retrieves all system cron jobs for the current user.
// It uses `sudo crontab -l` to list the cron jobs.
func allCrons() ([]string, error) {
	out, err := exec.Command("sudo", "crontab", "-l").Output()
	if err = cmdError(err); err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

// updateCrontab updates the crontab with the given content and restarts the cron service.
func updateCrontab(content string) error {
	// Update the crontab.
	cmd := `echo "` + content + `" | crontab -`
	if err := cmdError(exec.Command("sudo", "bash", "-c", cmd).Run()); err != nil {
		return err
	}

	// Restart the cron service to apply changes.
	return cmdError(exec.Command("sudo", "systemctl", "restart", "cron").Run())
}
