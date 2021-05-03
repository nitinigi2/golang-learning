package main

import (
	"fmt"
)

func main() {
	doSomething()
}

func doSomething() {
	slice := []int{1, 2, 3, 4, 5}
	modifyArray(&slice)
	fmt.Println(slice)
}

// slice pass by reference
func modifyArray(arr *[]int) {
	(*arr)[1] = 10
	*arr = append(*arr, 20)
}
