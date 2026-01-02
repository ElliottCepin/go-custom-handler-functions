package main

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T ) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	serveHello(rec, req)
	
	if (rec.Body.String() != "Hello, World!") {
		t.Errorf("Got %q, expected %q", rec.Body.String(),  "Hello, World!")
	}
}

func TestTime(t *testing.T) {
	req := httptest.NewRequest("GET", "/time", nil)
	rec := httptest.NewRecorder()
	
	serveTime(rec, req)

	// I don't know if it is possible to accurately test for correctness on the time...
	if (rec.Code != 200) {
		t.Errorf("Got code %q, expected code %q", rec.Code, 200)
	}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	serveHealth(rec, req)

	if (rec.Body.String() != "OK") {
		t.Errorf("Got %q, expected %q", rec.Body.String(), "OK")
	}
}
