package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	RegisterHandlers(r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
