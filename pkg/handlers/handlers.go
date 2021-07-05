package handlers

import (
	"github.com/ghadeerhamed/bookings/pkg/config"
	"github.com/ghadeerhamed/bookings/pkg/models"
	"github.com/ghadeerhamed/bookings/pkg/render"
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
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["test"] = "Hello!"
	stringMap["remote_ip"] = remoteIp

	//How to get the session in handlers
	//m.App.Session

	render.RenderTemplate(w, "about.page.tmpl",  &models.TemplateData{
		StringMap: stringMap,
	})
}
