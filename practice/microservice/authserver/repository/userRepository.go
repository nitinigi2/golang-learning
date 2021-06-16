package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bookApiDocker/authserver/db"
	"github.com/bookApiDocker/authserver/entity"
)

func IsUserValid(user *entity.User) (*entity.User, error) {
	conn := db.GetDBConnection()

	// close the db connection
	defer db.CloseDBConnection(conn)

	sqlStatement := `select * from USER where username = ? and password = ?`

	// execute the sql statement
	row := conn.QueryRow(sqlStatement, user.UserName, user.Password)

	err := row.Scan(&user.UserName, &user.Password)

	switch err {

	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return &entity.User{}, errors.New("no user found")

	case nil:
		return user, nil

	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, nil
}
