package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const httpHOST = "http://godoc.org"

var url string
var cachePage = cache.New(5*time.Minute, 10*time.Minute)

func main() {

	http.HandleFunc("/github.com/stretchr/testify/assert", loadPage)
	http.ListenAndServe(":8080", nil)

}

func loadPage(w http.ResponseWriter, r *http.Request) {

	url = httpHOST + r.URL.Path

	cachedResponse, found := cachePage.Get(r.URL.Path)

	if found {
		fmt.Println("Loading from cache")
		fmt.Fprintf(w, cachedResponse.(string))
	} else {
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		str := string(bodyBytes)
		cachePage.Set(r.URL.Path, str, cache.DefaultExpiration)

		fmt.Fprintf(w, str)
	}
}
