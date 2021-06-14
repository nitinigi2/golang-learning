package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nitinigi2/practice/book-api/model"
	"github.com/nitinigi2/practice/book-api/repository"
)

func SaveBook(w http.ResponseWriter, r *http.Request) {

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
}

func DeleteBook(w http.ResponseWriter, bookId int) {

	err := repository.DeleteBook(bookId)
	//err := service.DeleteBook(bookId)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, bookId int) {

	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//err := service.UpdateBook(book)
	err = repository.UpdateBook(book, bookId)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetBook(w http.ResponseWriter, bookId int) {

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
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := repository.GetAllBooks()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
