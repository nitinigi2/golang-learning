package service

import "github.com/nitinigi2/go-learning/testing/entity"

func add(x, y int) int {
	return x + y
}

func GetBooks() []entity.Book {
	books := make([]entity.Book, 0)

	books = append(books, entity.Book{
		Name:   "book1",
		Author: "author1",
	})

	books = append(books, entity.Book{
		Name:   "book2",
		Author: "author2",
	})

	return books
}
