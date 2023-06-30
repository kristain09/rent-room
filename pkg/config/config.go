package config

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func (a *AppConfig) InitConfig() {
	a.InProduction = false
	a.UseCache = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = a.InProduction
	a.Session = session
}
