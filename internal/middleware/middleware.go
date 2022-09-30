package middleware

import (
	"fmt"
	"net/http"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var App config.AppConfig
var Session *scs.SessionManager

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware triggered")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return Session.LoadAndSave(next)
}
