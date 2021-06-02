package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/rest-api/database"
	"github.com/nitinigi2/practice/rest-api/model"
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/books", booksHandler)
	r.HandleFunc("/api/books/{id}", bookHandler)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books, err := database.GetAllBooks()
		if err != nil {
			fmt.Println(err)
		}
		//books := service.GetBooks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var book model.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = database.SaveBook(book)
		//err = service.CreateBook(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Id not valid")
	}

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		book, err := database.GetBook(bookId)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(book)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	case http.MethodDelete:

		err := database.DeleteBook(bookId)
		//err := service.DeleteBook(bookId)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodPut:
		var book model.Book
		err = json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//err := service.UpdateBook(book)
		err := database.UpdateBook(book)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
