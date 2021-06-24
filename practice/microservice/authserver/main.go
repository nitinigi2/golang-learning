package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	RegisterRoutes(r)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
	}

	fmt.Println("listening on port", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("auth server couldn't start on port: ", srv.Addr)
	}

}
