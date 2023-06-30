package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	TempalateCache map[string]*template.Template
}
