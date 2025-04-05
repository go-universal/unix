package systemd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/go-universal/unix"
)

// SystemdService represents a systemd service manager.
type SystemdService interface {
	// Exists checks if the service exists.
	Exists() bool

	// Enabled checks if the service exists and enabled on startup.
	Enabled() bool

	// Disable disables the service.
	Disable() error

	// Install installs the service.
	// If override is false and the service already exists, it returns false.
	Install(override bool) (bool, error)

	// Uninstall uninstalls the service.
	Uninstall() error
}

// systemd is the implementation of the SystemdService interface.
type systemd struct {
	name string
	opt  *option
}

// NewService creates a new SystemdService instance with the given name, root, command and options.
func NewService(name, root, command string, options ...Option) SystemdService {
	name = strings.TrimSpace(name)
	root = strings.TrimSpace(root)
	command = strings.TrimSpace(command)
	engine := unix.NewTemplate().
		SetTemplate(serviceTemplate).
		AddParameter("name", name).
		AddParameter("root", root).
		AddParameter("command", command)

	option := &option{template: engine}
	for _, opt := range options {
		opt(option)
	}

	return &systemd{
		name: name,
		opt:  option,
	}
}

func (s *systemd) path() string {
	return "/etc/systemd/system/" + s.name + ".service"
}

func (s *systemd) Exists() bool {
	_, err := exec.Command("sudo", "systemctl", "status", s.name).Output()
	return err == nil
}

func (s *systemd) Enabled() bool {
	output, _ := exec.Command("sudo", "systemctl", "is-enabled", s.name).Output()
	return strings.HasPrefix(string(output), "enabled")
}

func (s *systemd) Disable() error {
	if s.Exists() {
		err := cmdError(exec.Command("sudo", "systemctl", "stop", s.name).Run())
		if err != nil {
			return err
		}

		err = cmdError(exec.Command("sudo", "systemctl", "disable", s.name).Run())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *systemd) Install(override bool) (bool, error) {
	if exists := s.Exists(); exists && !override {
		return false, nil
	}

	content := []byte(s.opt.template.Compile())
	if err := os.WriteFile(s.path(), content, 0644); err != nil {
		return false, err
	}

	if err := reload(); err != nil {
		return false, err
	}

	if err := cmdError(exec.Command("sudo", "systemctl", "enable", s.name).Run()); err != nil {
		return false, err
	}

	if err := cmdError(exec.Command("sudo", "systemctl", "start", s.name).Run()); err != nil {
		return false, err
	}

	return true, nil
}

func (s *systemd) Uninstall() error {
	if err := s.Disable(); err != nil {
		return err
	}

	if err := os.Remove(s.path()); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}
