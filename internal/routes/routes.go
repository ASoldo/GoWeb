package routes

import (
	"net/http"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/handlers"
	middlewares "github.com/ASoldo/GoWeb/internal/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.WriteToConsole)
	mux.Use(middlewares.NoSurf)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/reservations", handlers.Repo.Reservations)
	// mux.Post("/reservations", handlers.Repo.PostReservations)

	mux.Post("/reservations", handlers.Repo.PostitReservations)

	mux.Get("/getjson", handlers.Repo.JsonRq)
	mux.Post("/postjson", handlers.Repo.JsonPost)

	mux.Get("/reservations-summary", handlers.Repo.GetReservationsSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
