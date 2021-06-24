package main

import (
	"github.com/bookApiDocker/authserver/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", controller.Login).Methods("Post")
	r.HandleFunc("/logout", controller.Logout).Methods("Post")
}
