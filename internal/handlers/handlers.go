package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/ASoldo/GoWeb/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, my name is Andrija Hebrang"
	// w.Header().Set("Cache-Control", "max-age=2592000")
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{})
}
func (m *Repository) PostReservations(w http.ResponseWriter, r *http.Request) {
	datestring := r.Form.Get("inp")
	if len(datestring) == 0 {
		w.Write([]byte("Please enter a valid dates"))
		return
	}
	dates := strings.Split(datestring, ",")

	for i := range dates {
		fmt.Println(dates[i])
	}
	w.Write([]byte(fmt.Sprintf("Posted dates are: %s - %s", dates[0], dates[1])))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) JsonRq(w http.ResponseWriter, r *http.Request) {
	respo := jsonResponse{
		OK:      true,
		Message: "Sollllldo",
	}
	out, err := json.MarshalIndent(respo, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) JsonPost(w http.ResponseWriter, r *http.Request) {
	respo := jsonResponse{
		OK:      false,
		Message: "Sollllldo",
	}
	out, err := json.MarshalIndent(respo, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}