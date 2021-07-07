package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ghadeerhamed/bookings/internal/config"
	"github.com/ghadeerhamed/bookings/internal/models"
	"github.com/ghadeerhamed/bookings/internal/render"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["test"] = "Hello!"
	stringMap["remote_ip"] = remoteIp

	//How to get the session in handlers
	//m.App.Session

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) GeneralRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "general-room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MajorRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "major-room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form["start"]
	end := r.Form["end"]
	w.Write([]byte(fmt.Sprintf("Searching.... %v ----- %s", start, end)))
}

func (m *Repository) MakeReservationForm(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{OK: true, Message: "Available"}
	out, _ := json.MarshalIndent(resp, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
