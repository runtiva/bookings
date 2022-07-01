package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/runtiva/bookings/internal/config"
	"github.com/runtiva/bookings/internal/models"
	"github.com/runtiva/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	
	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP


	render.RenderTemplate(w, r, "about", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability", &models.TemplateData{})
}

type jsonResponse struct {
	OK		bool	`json:"ok"`
	Message	string	`json:"message"`
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse {
		OK: true,
		Message: "Available!",
	}
	
	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// PostAvailabilityJSON process the search availability JSON request
func (m *Repository) PostAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	var body []byte 
	n1, err1 := r.Body.Read(body)
	log.Println(n1, err1)
	log.Println(fmt.Sprintf("body: %s", string(body)))
	response := jsonResponse {
		OK: true,
		Message: "Available!",
	}
	
	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	// log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}


// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact", &models.TemplateData{})
}
