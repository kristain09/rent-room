package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kristain09/rent-room/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config fot the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the template cache from the app config

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get tempate from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Errror writing template for browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
