package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ASoldo/go-web/pkg/config"
	"github.com/ASoldo/go-web/pkg/handlers"
	"github.com/ASoldo/go-web/pkg/render"

	"runtime/trace"
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
	fmt.Println("Listening on port: 8080")
	tracer()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func tracer() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Tracing")
}
