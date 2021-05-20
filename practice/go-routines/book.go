package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprint("ID: ", b.ID, " Title: ", b.Title, " Author: ", b.Author)
}

func generateBookData() []Book {
	books := make([]Book, 0)
	tempId, tempTitle, tempAuthor := 0, "title", "author"
	for i := 1; i <= 10; i++ {
		var book Book = Book{
			ID:     tempId + i,
			Title:  tempTitle + fmt.Sprint(i),
			Author: tempAuthor + fmt.Sprint(i),
		}
		books = append(books, book)
	}
	return books
}

var Books = generateBookData()
