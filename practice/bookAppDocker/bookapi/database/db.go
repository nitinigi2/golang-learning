package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nitinigi2/practice/book-api/utilities"
)

var myEnv map[string]string = utilities.LoadEnvVaribles()
var DB *sql.DB

var (
	DB_DRIVER   = myEnv["DB_DRIVER"]
	DB_USER     = myEnv["DB_USER"]
	DB_PASSWORD = myEnv["DB_PASSWORD"]
	DB_NAME     = myEnv["DB_NAME"]
	DB_SERVER   = myEnv["DB_SERVER"]
	DB_PORT     = myEnv["DB_PORT"]
)

func GetDbConn() *sql.DB {

	connectDB()
	// return the connection
	return DB
}

func connectDB() {
	fmt.Println(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":"+DB_PORT+")/"+DB_NAME)
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

	DB = db

	log.Println("Successfully connected with database!")
}

func CloseDBConnection(db *sql.DB) {
	db.Close()
	log.Println("Successfully closed DB connection")
}
