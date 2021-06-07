package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getBook() {
	books := make([]Book, 0)

	bytesData := MakeRequest("GET", "/api/books", nil)

	json.Unmarshal(bytesData, &books)

	for _, book := range books {
		book.print()
		fmt.Println()
	}
}

func saveBook() {
	sampleBook := Book{
		Isbn:        "avdd",
		Title:       "ascd",
		Category:    "B category",
		Description: "cedsve",
		Author:      "frere",
	}

	bookJson, err := json.Marshal(sampleBook)

	handleError(err)

	fmt.Println(string(bookJson))

	r := bytes.NewReader(bookJson)

	resp, _ := http.Post(URL+"/api/books", "application/json", r)

	fmt.Println(resp.Status)
}

func updateBook() {

	sampleBook := Book{
		ID:          3,
		Isbn:        "avdd",
		Title:       "ascd",
		Category:    "B category",
		Description: "cedsve",
		Author:      "frere",
	}

	bookJson, err := json.Marshal(sampleBook)

	handleError(err)

	fmt.Println(string(bookJson))

	r := bytes.NewReader(bookJson)

	MakeRequest("PUT", "/api/books/"+fmt.Sprint(sampleBook.ID), r)
}

func DeleteBook(id int) {

	bytes := MakeRequest("DELETE", "/api/books/"+fmt.Sprint(id), nil)

	fmt.Println(string(bytes))
}

func handleError(err error) {
	if err != nil {
		log.Fatal("error ", err)
	}
}
