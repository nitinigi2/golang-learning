package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nitinigi2/practice/rest-api/configuration"
	"github.com/nitinigi2/practice/rest-api/model"
)

func SaveBook(book model.Book) error {
	// get db connection
	db := configuration.GetDbConn()
	fmt.Println("book", book)
	// close the db connection
	defer db.Close()

	sqlStatement := `INSERT INTO BOOK (isbn, title, category, description, author) VALUES (?, ?, ?, ?, ?)`

	// execute the sql statement
	// Scan function will save the insert id in the id

	_, err := db.Exec(sqlStatement, book.Isbn, book.Title, book.Category, book.Description, book.Author)

	if err != nil {
		fmt.Println("Unable to execute the query", err)
		return err
	}

	fmt.Println("Inserted a single record")
	return nil
}

func DeleteBook(id int) error {
	// get db connection
	db := configuration.GetDbConn()

	// close the db connection
	defer db.Close()

	sqlStatement := `delete from BOOK where id = ?`

	// execute the sql statement
	_, err := db.Exec(sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(book model.Book) error {
	// get db connection
	db := configuration.GetDbConn()

	// close the db connection
	defer db.Close()

	sqlStatement := `UPDATE BOOK SET isbn=?, title=?, category=?, description=?, author=? WHERE id=?`

	// execute the sql statement
	_, err := db.Exec(sqlStatement, book.Isbn, book.Title, book.Category, book.Description, book.Author, book.ID)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return err
}

func GetBook(id int) (model.Book, error) {
	// get db connection
	db := configuration.GetDbConn()

	// close the db connection
	defer db.Close()

	sqlStatement := `select * from BOOK where id = ?`

	var book model.Book
	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&book.ID, &book.Isbn, &book.Title, &book.Category, &book.Description, &book.Author)

	switch err {

	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return book, nil

	case nil:
		return book, nil

	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return book, err
}

func GetAllBooks() ([]model.Book, error) {
	// get db connection
	db := configuration.GetDbConn()

	// close the db connection
	defer db.Close()

	sqlStatement := `select * from BOOK`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	books := []model.Book{}

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Isbn, &book.Title, &book.Category, &book.Description, &book.Author)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
