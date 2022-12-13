package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ASoldo/GoWeb/internal/driver"
	"github.com/ASoldo/GoWeb/internal/handlers"
	"github.com/ASoldo/GoWeb/internal/helpers"
	"github.com/ASoldo/GoWeb/internal/middleware"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/ASoldo/GoWeb/internal/render"
	"github.com/ASoldo/GoWeb/internal/routes"
	"github.com/alexedwards/scs/v2"

	"runtime/trace"
)

func main() {
	//what am i going to put in session
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	db.SQL.Close()
	tracer()
	mystrl := "a,b,c,d"
	splitted := strings.Split(mystrl, ",")
	for i := range splitted {
		fmt.Println(splitted[i])
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(&middleware.App),
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// This is for profiling
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

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	middleware.App.InProduction = false

	middleware.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	middleware.App.InfoLog = middleware.InfoLog

	middleware.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	middleware.App.ErrorLog = middleware.ErrorLog

	middleware.Session = scs.New()
	middleware.Session.Lifetime = 24 * time.Hour
	middleware.Session.Cookie.Persist = true
	middleware.Session.Cookie.SameSite = http.SameSiteLaxMode
	middleware.Session.Cookie.Secure = middleware.App.InProduction
	middleware.App.Session = middleware.Session

	// conect to database
	log.Println("Connecting to the database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=test123")
	if err != nil {
		fmt.Println("Cannot connect to database! Exiting...")
	}
	fmt.Println("Connected to a database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache ", err)
		return nil, err
	}

	middleware.App.TemplateCache = tc
	middleware.App.UseCache = false
	repo := handlers.NewRepository(&middleware.App, db)
	handlers.NewHandlers(repo)
	render.NewTemplate(&middleware.App)
	helpers.NewHelpers(&middleware.App)
	fmt.Println("Server started")
	fmt.Println("Listening on port: 8080")
	return db, nil
}
