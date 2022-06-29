package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	InProduction    bool
	UseCache 	    bool
	TemplateCache   map[string]*template.Template
	InfoLog         *log.Logger	
	SessionManager  *scs.SessionManager
}