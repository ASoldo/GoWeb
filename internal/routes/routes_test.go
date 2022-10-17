package routes

import (
	"fmt"
	"testing"

	"github.com/ASoldo/GoWeb/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := Routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error("Type is not *chi.Mux")
		fmt.Printf("Type is %T", v)
	}
}
