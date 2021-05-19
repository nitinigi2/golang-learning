package main

import "fmt"

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	fmt.Println(rev)
	return temp == rev
}

func main() {
	x := 121

	fmt.Println(isPalindrome(x))
}
