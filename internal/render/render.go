package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/runtiva/bookings/internal/config"
	"github.com/runtiva/bookings/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
// NewTemplates sets config for template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Renders HTML Templates by name
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatalln("Unable to Create Template Cache during runtime")
		}
	}

	t, ok := tc[tmpl+".page.go.tmpl"]
	if !ok {
		log.Fatal("template not found in cache")
	}
	
	buf := new(bytes.Buffer)
	
	// Add default data to the model 
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	cnt, err := buf.WriteTo(w)
	if cnt < 1 {
		fmt.Println("WriteTo error", cnt)
	}
	if err != nil {
		fmt.Println("Error writing template to browser", err) 
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.go.html")
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error parsing template:", err)
	// }
}

// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.go.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//name = name[:strings.IndexByte(name, '.')]
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.go.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.go.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
		log.Println(name)
	}

	return myCache, nil
}