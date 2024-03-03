package enthistory

import (
	"embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed templates/*
	_templates embed.FS
)

type Config struct {
	Auditing  bool
	UpdatedBy *UpdatedBy
}

func (c Config) Name() string {
	return "HistoryConfig"
}

// HistoryExtension implements entc.Extension.
type HistoryExtension struct {
	entc.DefaultExtension
	config *Config
}

type ExtensionOption = func(*HistoryExtension)

// WithAuditing allows you to turn on the code generation for the `.Audit()` method
func WithAuditing() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.Auditing = true
	}
}

func NewHistoryExtension(opts ...ExtensionOption) *HistoryExtension {
	extension := &HistoryExtension{
		// Set configuration defaults that can get overridden with ExtensionOption
		config: &Config{
			Auditing:  false,
			UpdatedBy: updatedBy,
		},
	}
	for _, opt := range opts {
		opt(extension)
	}

	return extension
}

func (h *HistoryExtension) Templates() []*gen.Template {
	templates := []*gen.Template{
		parseTemplate("historyFromMutation", "templates/historyFromMutation.tmpl"),
		parseTemplate("historyQuery", "templates/historyQuery.tmpl"),
		parseTemplate("client", "templates/client.tmpl"),
	}
	if h.config.Auditing {
		templates = append(templates, parseTemplate("auditing", "templates/auditing.tmpl"))
	}
	return templates
}

// Hooks of the HistoryExtension.
func (h *HistoryExtension) Hooks() []gen.Hook {
	return []gen.Hook{}
}

func (h *HistoryExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		h.config,
	}
}
