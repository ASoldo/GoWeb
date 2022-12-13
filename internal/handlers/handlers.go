package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/driver"
	"github.com/ASoldo/GoWeb/internal/forms"
	"github.com/ASoldo/GoWeb/internal/helpers"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/ASoldo/GoWeb/internal/render"
	"github.com/ASoldo/GoWeb/internal/repository"
	"github.com/ASoldo/GoWeb/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepository() creates a new repository
func NewRepository(ac *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: ac,
		DB:  dbrepo.NewPostgresRepo(db.SQL, ac),
	}
}

// NewHandlers() sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home() is the Home Page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About() is the About Page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, my name is Andrija Hebrang"
	// w.Header().Set("Cache-Control", "max-age=2592000")
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservations() is the Reservations Page handler
func (m *Repository) Reservations(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservations() is the PostReservations Page handler
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
	// err = errors.New("this is an error message soldo")
	if err != nil {
		// fmt.Println(err)
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		StartEndDate:    r.Form.Get("inp"),
		AdditionalInput: r.Form.Get("additionalInput"),
	}
	form := forms.New(r.PostForm)

	// form.Has("inp", r)
	form.Required("inp", "additionalInput")
	form.MinLength("additionalInput", 3)
	// form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "reservations.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservations-summary", http.StatusSeeOther)
}

func (m *Repository) GetReservationsSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation) //type cast to models.Reservation
	if !ok {
		// fmt.Println("Cannot get item from session")
		m.App.ErrorLog.Println(w, ok) // can't get error from session
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from this session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//remove 'reservation' from session
	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservations-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
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
		// fmt.Println(err)
		helpers.ServerError(w, err)
		return
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
