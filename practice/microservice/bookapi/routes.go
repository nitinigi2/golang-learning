package main

import (
	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/book-api/controller"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/books", controller.BooksHandler)
	r.HandleFunc("/api/books/{id:[0-9]+}", controller.BookHandler)
}
