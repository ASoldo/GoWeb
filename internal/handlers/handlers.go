package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/forms"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
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

// PostitReservations handlers the posting of a reservation form
func (m *Repository) PostitReservations(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	reservation := models.Reservation{
		StartEndDate:    r.Form.Get("inp"),
		AdditionalInput: r.Form.Get("additionalInput"),
	}
	form := forms.New(r.PostForm)

	// form.Has("inp", r)
	form.Required("inp", "additionalInput")
	form.MinLength("additionalInput", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

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
	datastring := strings.Split(r.Form.Get("inp"), ",")
	data := datastring
	for i := range data {
		fmt.Println(data[i])
	}
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
