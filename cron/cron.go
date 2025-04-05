package cron

import (
	"strings"
)

// Cron represents a scheduled job manager.
type Cron interface {
	// Raw returns the raw cron expression.
	Raw() string

	// Exists checks whether the cron job is installed.
	Exists() (bool, error)

	// Install sets up the cron job.
	// Returns false if it already exists.
	Install() (bool, error)

	// Uninstall removes the cron job.
	Uninstall() error
}

// cron is the implementation of the Cron interface.
type cron struct {
	opt     *option
	command string
}

// New creates a new Cron instance with the given command and options.
func New(command string, options ...Option) Cron {
	option := &option{
		tz:      NewTZ(),
		reboot:  false,
		minute:  "*",
		hour:    "*",
		day:     "*",
		month:   "*",
		weekday: "*",
	}
	for _, opt := range options {
		opt(option)
	}

	return &cron{
		opt:     option,
		command: command,
	}
}

func (c *cron) Raw() string {
	if c.opt.reboot {
		return "@reboot " + c.command
	}
	return c.opt.interval() + " " + c.command
}

func (c *cron) Exists() (bool, error) {
	lines, err := allCrons()
	if err != nil {
		return false, err
	}

	for _, line := range lines {
		ok, cmd := parseCommand(line)
		if ok && cmd == c.command {
			return true, nil
		}
	}

	return false, nil
}

func (c *cron) Install() (bool, error) {
	var exists bool
	var result strings.Builder

	lines, err := allCrons()
	if err != nil {
		return false, err
	}

	for _, line := range lines {
		ok, cmd := parseCommand(line)
		if ok && cmd == c.command {
			exists = true
			result.WriteString(c.Raw() + "\n")
			continue
		}

		if strings.TrimSpace(line) != "" {
			result.WriteString(line + "\n")
		}
	}

	if !exists {
		result.WriteString(c.Raw() + "\n")
	}

	if err := updateCrontab(result.String()); err != nil {
		return false, err
	}

	return true, nil
}

func (c *cron) Uninstall() error {
	var result strings.Builder

	lines, err := allCrons()
	if err != nil {
		return err
	}

	for _, line := range lines {
		ok, cmd := parseCommand(line)
		if ok && cmd == c.command {
			continue
		}

		if strings.TrimSpace(line) != "" {
			result.WriteString(line + "\n")
		}
	}

	return updateCrontab(result.String())
}
