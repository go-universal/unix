package nginx

import (
	"os"
	"strings"

	"github.com/go-universal/unix"
)

// ServerBlock represents a NGINX server block manager.
type ServerBlock interface {
	// Exists returns whether the site configuration exists.
	Exists() (bool, error)

	// Enabled returns whether the site exists and is currently enabled.
	Enabled() (bool, error)

	// Disable disables the site.
	Disable() error

	// Enable enables the site.
	Enable() error

	// Install sets up the site configuration.
	// If override is false and the site already exists, it returns false.
	Install(override bool) (bool, error)

	// Uninstall removes the site configuration.
	Uninstall() error
}

// serverBlock is the implementation of the ServerBlock interface.
type serverBlock struct {
	name string
	opt  *option
}

// NewServerBlock creates a new ServerBlock instance with the given name, template and options.
func NewServerBlock(name, template string, options ...Option) ServerBlock {
	name = strings.TrimSpace(name)
	engine := unix.NewTemplate().SetTemplate(template)

	option := &option{template: engine}
	for _, opt := range options {
		opt(option)
	}

	return &serverBlock{
		name: name,
		opt:  option,
	}
}

func (s *serverBlock) path() string {
	return "/etc/nginx/sites-available/" + s.name
}

func (s *serverBlock) link() string {
	return "/etc/nginx/sites-enabled/" + s.name
}

func (s *serverBlock) Exists() (bool, error) {
	return fileExists(s.path())
}

func (s *serverBlock) Enabled() (bool, error) {
	available, err := fileExists(s.path())
	if err != nil {
		return false, err
	}

	enabled, err := fileExists(s.link())
	if err != nil {
		return false, err
	}

	return available && enabled, nil
}

func (s *serverBlock) Disable() error {
	if err := os.Remove(s.link()); err != nil && !os.IsNotExist(err) {
		return err
	}

	return restart()
}

func (s *serverBlock) Enable() error {
	exists, err := fileExists(s.link())
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if err := os.Symlink(s.path(), s.link()); err != nil {
		return err
	}

	return restart()
}

func (s *serverBlock) Install(override bool) (bool, error) {
	exists, err := fileExists(s.path())
	if err != nil {
		return false, err
	}

	if exists && !override {
		return false, nil
	}

	content := []byte(s.opt.template.Compile())
	if err := os.WriteFile(s.path(), content, 0644); err != nil {
		return false, err
	}

	if err := s.Enable(); err != nil {
		return false, err
	}

	return true, nil
}

func (s *serverBlock) Uninstall() error {
	if err := os.Remove(s.link()); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := os.Remove(s.path()); err != nil && !os.IsNotExist(err) {
		return err
	}

	return restart()
}
