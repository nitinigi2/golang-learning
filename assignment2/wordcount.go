// https://www.hackerrank.com/contests/codart-2-0/challenges/word-count-1/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var noOfString int
	_, err := fmt.Scanf("%d\n", &noOfString)

	checkError(err)
	reader := bufio.NewReader(os.Stdin)
	//reader.ReadString('\n')
	dict, lineNo := make(map[string][]int), 1
	for noOfString > 0 {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " ")
		populateMap(text, dict, lineNo)
		lineNo++
		noOfString--
	}
	printMap(dict)
}

func printMap(dict map[string][]int) {
	slice := sortKeys(dict)
	for _, value := range slice {
		fmt.Print(value)
		printSlice(dict[value])
	}
}

func printSlice(s []int) {
	for _, value := range s {
		fmt.Print(" ", value)
	}
	fmt.Println()
}

func sortKeys(dict map[string][]int) []string {
	slice, i := make([]string, len(dict)), 0

	for key, _ := range dict {
		slice[i] = key
		i++
	}

	sort.Strings(slice)
	return slice
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func populateMap(text string, dict map[string][]int, lineNo int) {
	words := strings.Fields(text)

	for _, value := range words {
		if _, ok := dict[value]; ok {
			dict[value] = append(dict[value], lineNo)
		} else {
			dict[value] = []int{lineNo}
		}
	}
}
