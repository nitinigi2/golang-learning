package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER   = "mysql"
	DB_USER     = "7MBlm8Z2XX"
	DB_PASSWORD = "GlDhD4mSVn"
	DB_NAME     = "7MBlm8Z2XX"
	DB_SERVER   = "remotemysql.com"
	DB_PORT     = "3306"
)

func GetDbConn() *sql.DB {

	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":"+DB_PORT+")/"+DB_NAME)

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
