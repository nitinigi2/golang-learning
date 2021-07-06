package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/go-learning/testing/controller"
)

func main() {
	registerRoutes()
	startServer()
}

func startServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		srv.ListenAndServe()
		wg.Done()
	}()

	return srv
}

func registerRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")
	return r
}
