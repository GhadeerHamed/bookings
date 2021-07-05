package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ghadeerhamed/bookings/pkg/config"
	"github.com/ghadeerhamed/bookings/pkg/handlers"
	"github.com/ghadeerhamed/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const addr = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//In Production Mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error Creating template cache: ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	fmt.Printf("App started on: %v\n\n", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
