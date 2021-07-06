package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	srv := startServer()
	srv.Shutdown(context.TODO())
}

func TestRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/books", nil)
	response := httptest.NewRecorder()
	registerRoutes().ServeHTTP(response, req)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
