package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	content := "Hello from go"

	file, err := os.Create("/.fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("No of char written to file", ln)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() {
	filename := "./hello.txt"
	content, err := ioutil.ReadFile(filename)
	checkError(err)
	result := string(content) // convert byte array to string
	fmt.Println("Read from file: ", result)
}

func readFileFromWeb() {
	// modify url with some url that gives json response
	url := "sample path"

	content := getContentFromWeb(url)

	fmt.Println(content)
}

func getContentFromWeb(url string) string {
	resp, err := http.Get(url)

	checkError(err)
	fmt.Println(resp)
	// ensure response body is closed after reading
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	checkError(err)

	return string(bytes)
}
