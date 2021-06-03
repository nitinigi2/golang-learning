package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const URL = "http://localhost:8080"

func MakeRequest(reqType string, path string, body io.Reader) []byte {
	req, err := http.NewRequest(reqType, URL+path, body)

	req.Header.Set("Content-Type", "application/json")

	fmt.Println("request URL: ", req.URL)
	fmt.Println("request method: ", req.Method)

	handleError(err)

	res, err1 := client.Do(req)

	fmt.Println("response status: ", res.Status)

	handleError(err1)

	defer res.Body.Close()

	byteData, err := ioutil.ReadAll(res.Body)

	handleError(err)

	return byteData
}

func getClient() *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	return client
}
