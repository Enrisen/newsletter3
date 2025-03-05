package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}
	req := httptest.NewRequest("GET", "/home", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.home)
	handler.ServeHTTP(rr, req)
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v, expected %v", status, http.StatusOK)
	}
	expected := "this the sign up page"
	got := rr.Body.String()
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
