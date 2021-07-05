package main

import (
	"github.com/ghadeerhamed/bookings/pkg/config"
	"github.com/ghadeerhamed/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/rooms/generals-quarters", handlers.Repo.GeneralRoom)
	mux.Get("/rooms/major-suit", handlers.Repo.MajorRoom)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/make-reservation-form", handlers.Repo.MakeReservationForm)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
