package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ghadeerhamed/bookings/internal/config"
	"github.com/ghadeerhamed/bookings/internal/driver"
	"github.com/ghadeerhamed/bookings/internal/handlers"
	"github.com/ghadeerhamed/bookings/internal/helpers"
	"github.com/ghadeerhamed/bookings/internal/models"
	"github.com/ghadeerhamed/bookings/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const addr = ":8000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	//Close connection after app finish
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {
			_ = fmt.Errorf("error Disconnecting with database [%v]", err)
		}
	}(db.SQL)

	fmt.Printf("App started on: %v\n\n", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	//What I am going to put in the session
	gob.Register(models.Reservation{})

	//In Production Mode
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//Connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=root")
	if err != nil {
		log.Fatal("Connot connect to database!. Dying ", err)
	}
	log.Println("Connected to database!")

	var tc map[string]*template.Template
	tc, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error Creating template cache: ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
