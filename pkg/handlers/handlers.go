package handlers

import (
	"net/http"

	"github.com/runtiva/bookings/pkg/config"
	"github.com/runtiva/bookings/pkg/models"
	"github.com/runtiva/bookings/pkg/render"
)

// Repo the repository is used by the handlers
var Repo *Repository

// Repository si the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHaldners sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	
	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP


	render.RenderTemplate(w, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}
