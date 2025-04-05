package systemd

import (
	"strings"

	"github.com/go-universal/unix"
)

// option holds the configuration for a systemd service.
type option struct {
	template unix.TemplateEngine
}

// Option defines a functional option for configuring settings.
type Option func(*option)

// WithTemplate sets the template string for the systemd service.
func WithTemplate(template string) Option {
	template = strings.TrimSpace(template)
	return func(o *option) {
		if template != "" {
			o.template.SetTemplate(template)
		}
	}
}

// WithParameter adds a parameter to replace in the template.
func WithParameter(name, value string) Option {
	name = strings.TrimSpace(name)
	return func(o *option) {
		if name != "" {
			o.template.AddParameter(name, value)
		}
	}
}
