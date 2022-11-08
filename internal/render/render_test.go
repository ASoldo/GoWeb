package render

import (
	"net/http"
	"testing"

	"github.com/ASoldo/GoWeb/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := GetSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "124")

	result := AddDefaultData(&td, r)
	if result.Flash != "124" {
		t.Error("flash value of 124 not found in session")
	}
}

func GetSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates/"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := GetSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	// err = RenderTemplate(&ww, r, "home-non.page.tmpl", &models.TemplateData{})
	// if err != nil {
	// 	t.Error("rendered template that does not exist")
	// }

}

func TestNewTemplates(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates/"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
