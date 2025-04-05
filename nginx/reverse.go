package nginx

import (
	"os"
	"strings"

	"github.com/go-universal/unix"
)

// ReverseProxy represents a NGINX reverse proxy manager.
type ReverseProxy interface {
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

// reverse is the implementation of the ReverseProxy interface.
type reverse struct {
	name string
	opt  *option
}

// NewReverseProxy creates a new ReverseProxy instance with the given name, port, domains and options.
func NewReverseProxy(name, port string, domains []string, options ...Option) ReverseProxy {
	name = strings.TrimSpace(name)
	port = strings.TrimSpace(port)
	trimmed := make([]string, 0, len(domains))
	for _, d := range domains {
		d = strings.TrimSpace(d)
		if d != "" {
			trimmed = append(trimmed, d)
		}
	}

	engine := unix.NewTemplate().
		SetTemplate(reverseTemplate).
		AddParameter("port", strings.TrimSpace(port)).
		AddParameter("domains", strings.Join(trimmed, " "))

	option := &option{template: engine}
	for _, opt := range options {
		opt(option)
	}

	return &reverse{
		name: name,
		opt:  option,
	}
}

func (r *reverse) path() string {
	return "/etc/nginx/sites-available/" + r.name
}

func (r *reverse) link() string {
	return "/etc/nginx/sites-enabled/" + r.name
}

func (r *reverse) Exists() (bool, error) {
	return fileExists(r.path())
}

func (r *reverse) Enabled() (bool, error) {
	available, err := fileExists(r.path())
	if err != nil {
		return false, err
	}

	enabled, err := fileExists(r.link())
	if err != nil {
		return false, err
	}

	return available && enabled, nil
}

func (r *reverse) Disable() error {
	if err := os.Remove(r.link()); err != nil && !os.IsNotExist(err) {
		return err
	}

	return restart()
}

func (r *reverse) Enable() error {
	exists, err := fileExists(r.link())
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if err := os.Symlink(r.path(), r.link()); err != nil {
		return err
	}

	return restart()
}

func (r *reverse) Install(override bool) (bool, error) {
	exists, err := fileExists(r.path())
	if err != nil {
		return false, err
	}

	if exists && !override {
		return false, nil
	}

	content := []byte(r.opt.template.Compile())
	if err := os.WriteFile(r.path(), content, 0644); err != nil {
		return false, err
	}

	if err := r.Enable(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *reverse) Uninstall() error {
	if err := os.Remove(r.link()); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := os.Remove(r.path()); err != nil && !os.IsNotExist(err) {
		return err
	}

	return restart()
}
