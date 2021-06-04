package main

import (
	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/rest-api/handler"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/books", handler.BooksHandler)
	r.HandleFunc("/api/books/{id}", handler.BookHandler)
}
