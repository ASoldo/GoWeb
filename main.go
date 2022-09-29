package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/handlers"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/ASoldo/GoWeb/internal/render"
	"github.com/alexedwards/scs/v2"

	"runtime/trace"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//what am i going to put in session
	gob.Register(models.Reservation{})
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
	mystrl := "a,b,c,d"
	splitted := strings.Split(mystrl, ",")
	for i := range splitted {
		fmt.Println(splitted[i])
	}

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
