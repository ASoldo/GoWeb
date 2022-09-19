package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ASoldo/go-web/pkg/config"
	"github.com/ASoldo/go-web/pkg/handlers"
	"github.com/ASoldo/go-web/pkg/render"
	"github.com/alexedwards/scs/v2"

	"runtime/trace"
)

var app config.AppConfig
var session *scs.SessionManager

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
//  1. Replace multiple slashes with a single slash.
//  2. Eliminate each . path name element (the current directory).
//  3. Eliminate each inner .. path name element (the parent directory)
//     along with the non-.. element that precedes it.
//  4. Eliminate .. elements that begin a rooted path:
//     that is, replace "/.." by "/" at the beginning of a path.
//
// The returned path ends in a slash only if it is the root "/".
//
// If the result of this process is an empty string, Clean
// returns the string ".".
//
// See also Rob Pike, “[Lexical File Names in Plan 9].”
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
