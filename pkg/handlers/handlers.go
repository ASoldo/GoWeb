package handlers

import (
	"net/http"

	"github.com/ASoldo/GoWeb/pkg/config"
	"github.com/ASoldo/GoWeb/pkg/models"
	"github.com/ASoldo/GoWeb/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, my name is Andrija Hebrang"
	// w.Header().Set("Cache-Control", "max-age=2592000")
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservations.page.tmpl", &models.TemplateData{})
}
