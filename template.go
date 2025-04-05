package unix

import "strings"

// TemplateEngine bracket wrapped string template with strings.Replacer.
type TemplateEngine interface {
	// Sets the template string.
	SetTemplate(template string) TemplateEngine

	// Adds a parameter to replace in the template.
	AddParameter(name, value string) TemplateEngine

	// Compiles the template by replacing placeholders with their values.
	Compile() string
}

// templateEngine is the implementation of the TemplateEngine interface.
type templateEngine struct {
	template string
	params   []string
}

// NewTemplate creates and initializes a new TemplateEngine instance.
func NewTemplate() TemplateEngine {
	engine := new(templateEngine)
	engine.params = make([]string, 0)
	return engine
}

func (t *templateEngine) SetTemplate(template string) TemplateEngine {
	t.template = template
	return t
}

func (t *templateEngine) AddParameter(name, value string) TemplateEngine {
	t.params = append(t.params, "{"+name+"}", value)
	return t
}

func (t *templateEngine) Compile() string {
	return strings.NewReplacer(t.params...).Replace(t.template)
}
