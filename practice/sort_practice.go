package main

import (
	"fmt"
	"sort"
)

func doSomething() {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Aichael", 31},
		{"Jenny", 26},
	}
	// sort based on age first
	// if age is same then sort based on name
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}

func main() {
	doSomething()
}
