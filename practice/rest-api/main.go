package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	r := mux.NewRouter()

	RegisterHandlers(r)

	server := readCofigFile()
	serverIp := server.Ip
	serverPort := server.Port

	log.Fatal(http.ListenAndServe(serverIp+":"+serverPort, r))

}

func readCofigFile() *Server {
	var server Server
	jsonFile, err := os.Open("serverconfig.json")
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &server)

	return &server
}
