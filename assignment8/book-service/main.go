package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/rest-api/handler"
)

func main() {
	r := mux.NewRouter()

	handler.RegisterHandlers(r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
