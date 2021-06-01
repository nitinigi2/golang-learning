package main

import (
	"database/sql"
	"fmt"
)

const (
	DB_DRIVER   = "mysql"
	DB_USER     = "root"
	DB_PASSWORD = "password"
	DB_NAME     = "db_name"
)

func GetDbConn() *sql.DB {

	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@/"+DB_NAME)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
