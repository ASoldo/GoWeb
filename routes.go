package main

import (
	"net/http"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/reservations", handlers.Repo.Reservations)
	mux.Post("/reservations", handlers.Repo.PostReservations)

	mux.Post("/reservations1", handlers.Repo.PostitReservations)

	mux.Get("/getjson", handlers.Repo.JsonRq)
	mux.Post("/postjson", handlers.Repo.JsonPost)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
