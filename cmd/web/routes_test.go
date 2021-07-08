package main

import (
	"fmt"
	"github.com/ghadeerhamed/bookings/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//Do nothing
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux, It is:%T", v))
	}
}
