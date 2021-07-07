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
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/make-reservation", handlers.Repo.MakeReservationForm)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/about", handlers.Repo.About)

	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
