// This file will do crud operations in-memory without any db
// it reads sample json data to load data initially
// to use this file call these crud methods from handler file
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nitinigi2/practice/rest-api/model"
)

// temp storage to store all books with their ids
var booksMap = make(map[int]model.Book)

var books = make([]model.Book, 0)

// keeping bookcounter, so new book id can be auto-increment
var bookCounter = 1

func init() {
	loadSampleJson()
	loadBooksMap()
	bookCounter = len(books)
	fmt.Println("Successfully loaded books.json")
}

// load sample json books file and store in slice of books
func loadSampleJson() {
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	//var books Books
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &books)
}

// load map with books id as key and book as value
func loadBooksMap() {
	for _, book := range books {
		booksMap[book.ID] = book
	}
}

// get api for all books
func GetBooks() []model.Book {
	books := make([]model.Book, 0)

	for _, book := range booksMap {
		books = append(books, book)
	}

	return books
}

// get api for single book
func GetBook(id int) (model.Book, error) {
	book, ok := booksMap[id]

	if ok {
		return book, nil
	}

	return model.Book{}, errors.New("book not found")
}

// post api for creating book
func CreateBook(book model.Book) error {
	if book.ID != 0 {
		return errors.New("book cannot be saved")
	}

	bookCounter++

	book.ID = bookCounter

	booksMap[book.ID] = book

	return nil
}

// put api for updating book
func UpdateBook(book model.Book) error {
	_, ok := booksMap[book.ID]

	if !ok {
		return errors.New("book not present")
	}

	booksMap[book.ID] = book
	return nil
}

// delete api for deleting book
func DeleteBook(id int) error {
	_, ok := booksMap[id]

	if !ok {
		return errors.New("book not present")
	}

	delete(booksMap, id)

	return nil
}
