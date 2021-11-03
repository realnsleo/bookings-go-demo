package handlers

import (
	"github.com/realnsleo/bookings-go-demo/pkg/config"
	"github.com/realnsleo/bookings-go-demo/pkg/models"
	"github.com/realnsleo/bookings-go-demo/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.BuildTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.BuildTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
