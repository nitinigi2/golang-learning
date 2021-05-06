// https://www.hackerrank.com/challenges/pangrams/problem
package main

import (
	"fmt"
	"strings"
)

func pangrams(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	fmt.Println(s)
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			return "not pangram"
		}
	}
	return "pangram"
}

func main() {
	s := "We promptly judged antique ivory buckles for the prize"
	fmt.Println(pangrams(s))
}
