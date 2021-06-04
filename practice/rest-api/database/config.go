package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	DB_DRIVER   string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SERVER   string
	DB_PORT     string
}

var db *DB

func init() {
	db = readCofigFile()
}

func GetDbConn() *sql.DB {
	db, err := sql.Open(db.DB_DRIVER, db.DB_USER+":"+db.DB_PASSWORD+"@tcp("+db.DB_SERVER+":"+db.DB_PORT+")/"+db.DB_NAME)

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

func readCofigFile() *DB {
	var db DB
	jsonFile, err := os.Open("dbconfig.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &db)

	fmt.Println(db)
	return &db
}
