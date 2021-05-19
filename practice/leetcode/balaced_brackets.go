package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	dict := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else if len(stack) == 0 || stack[len(stack)-1] != dict[ch] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
