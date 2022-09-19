package main

import (
	"fmt"
	"net/http"

	"github.com/ASoldo/go-web/pkg/config"
	"github.com/ASoldo/go-web/pkg/handlers"
	"github.com/ASoldo/go-web/pkg/render"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	fmt.Println("Server started")
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
