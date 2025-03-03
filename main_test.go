package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupHandler(t *testing.T) {
    // set up request and response
    req := httptest.NewRequest("GET", "/home", nil)
    rr := httptest.NewRecorder()
    // convert the home  function to a http.Handler type
    handler := http.HandlerFunc(home)
    // pass the fake HTTP request to the handler
    handler.ServeHTTP(rr, req)
    // perform  the assertions
    //1. check the status code for 200 OK
    status := rr.Code
    if status != http.StatusOK {
      t.Errorf("got %v, expected %v", status, http.StatusOK)
    }
    //2. check the response body
    expected := "this the sign up page"
    got := rr.Body.String()
    if got != expected {
      t.Errorf("got %v, expected %v", got, expected)
    }
}
