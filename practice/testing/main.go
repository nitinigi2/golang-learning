package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/go-learning/testing/controller"
)

func main() {
	r := mux.NewRouter()
	registerRoutes(r)
	startServer()
}

func startServer() {
	http.ListenAndServe(":8080", nil)
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")
}
