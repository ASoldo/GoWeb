package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/ASoldo/GoWeb/internal/middleware"
	"github.com/ASoldo/GoWeb/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	middleware.App.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	middleware.App.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	testApp.Session = session
	app = &testApp
	os.Exit(m.Run())
}

type myWriter struct{}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWriter) WriteHeader(i int) {

}

func (mw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
