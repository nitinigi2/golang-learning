package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	DB_DRIVER   = os.Getenv("DB_DRIVER")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_SERVER   = os.Getenv("DB_SERVER")
	DB_PORT     = os.Getenv("DB_PORT")
)

func connectDB() {

	db, err := sql.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@tcp("+DB_SERVER+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}

	DB = db

	log.Println("Successfully connected with database")
}

func GetDBConnection() *sql.DB {
	connectDB()
	return DB
}

func CloseDBConnection(db *sql.DB) {
	db.Close()
	log.Println("Successfully closed DB connection")
}
