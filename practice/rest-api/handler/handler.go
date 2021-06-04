package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinigi2/practice/rest-api/model"
	"github.com/nitinigi2/practice/rest-api/repository"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books, err := repository.GetAllBooks()
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
		err = repository.SaveBook(book)
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

func BookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	bookId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Id not valid")
	}

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		book, err := repository.GetBook(bookId)

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

		err := repository.DeleteBook(bookId)
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
		err := repository.UpdateBook(book)
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
