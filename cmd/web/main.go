package main

import (
	"log"
	"net/http"
	"time"

	"github.com/runtiva/bookings/pkg/config"
	"github.com/runtiva/bookings/pkg/handlers"
	"github.com/runtiva/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	
	// Initialize config and push into render
	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	portNumber := "127.0.0.1:8080"
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

