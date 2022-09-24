package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ASoldo/GoWeb/pkg/config"
	"github.com/ASoldo/GoWeb/pkg/handlers"
	"github.com/ASoldo/GoWeb/pkg/render"
	"github.com/alexedwards/scs/v2"

	"runtime/trace"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
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
	fmt.Println("Tracing started")
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

	fmt.Println("Tracing end")
}
