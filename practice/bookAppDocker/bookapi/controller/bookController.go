package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/book-api/service"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		service.GetBooks(w, r)

	case http.MethodPost:
		service.SaveBook(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Id not valid")
	}

	switch r.Method {

	case http.MethodGet:
		service.GetBook(w, bookId)

	case http.MethodDelete:
		service.DeleteBook(w, bookId)

	case http.MethodPut:
		service.UpdateBook(w, r, bookId)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
