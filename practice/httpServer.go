package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello from the Go web Server</h1>")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var h Hello

	err := http.ListenAndServe("localhost:8080", h)

	checkError(err)
}