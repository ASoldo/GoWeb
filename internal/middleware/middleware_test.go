package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)
	if h == nil {
		t.Error("We have some error here")
	}
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error("Not http.handler")
		fmt.Printf("%T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)
	if h == nil {
		t.Error("We have some error here")
	}
	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error("Not http.handler")
		fmt.Printf("%T", v)
	}
}
