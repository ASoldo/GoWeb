package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	middlewares "github.com/ASoldo/GoWeb/internal/middleware"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/ASoldo/GoWeb/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var pathToTemplates = "./../../templates/"

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})
	middlewares.App.InProduction = false
	middlewares.Session = scs.New()
	middlewares.Session.Lifetime = 24 * time.Hour
	middlewares.Session.Cookie.Persist = true
	middlewares.Session.Cookie.SameSite = http.SameSiteLaxMode
	middlewares.Session.Cookie.Secure = middlewares.App.InProduction
	middlewares.App.Session = middlewares.Session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache ", err)
		// return err
	}

	middlewares.App.TemplateCache = tc
	middlewares.App.UseCache = true
	repo := NewRepository(&middlewares.App)
	NewHandlers(repo)
	render.NewTemplate(&middlewares.App)
	fmt.Println("Server started")
	fmt.Println("Listening on port: 8080")
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.WriteToConsole)
	// mux.Use(middlewares.NoSurf)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)

	mux.Get("/reservations", Repo.Reservations)
	// mux.Post("/reservations", Repo.PostReservations)

	mux.Post("/reservations", Repo.PostitReservations)

	mux.Get("/getjson", Repo.JsonRq)
	mux.Post("/postjson", Repo.JsonPost)

	mux.Get("/reservations-summary", Repo.GetReservationsSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all of the names .page.tmpl files
	pages, err := filepath.Glob(fmt.Sprintf("%s*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s*.layout.tmpl", pathToTemplates))
		if err != nil {
			fmt.Println(err)
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
