package handlers

import (
	"net/http"

	"github.com/ASoldo/go-web/pkg/config"
	"github.com/ASoldo/go-web/pkg/models"
	"github.com/ASoldo/go-web/pkg/render"
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
	// fmt.Fprintf(w, "Welcome to the Internet")
	// fmt.Println(w)
	// fmt.Println(r)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is About page")
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, my name is Andrija Hebrang"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
