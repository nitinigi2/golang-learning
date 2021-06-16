package main

import (
	"fmt"
	"net/http"

	"github.com/bookApiDocker/authserver/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.Login).Methods("Post")
	r.HandleFunc("/logout", controller.Logout).Methods("Post")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("listening on port 8080")
	srv.ListenAndServe()
}
