package main

import "fmt"

type Book struct {
	ID          int    `json:"id"`
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (book Book) print() {
	fmt.Printf("%+v", book)
}
