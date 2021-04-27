package main

import "fmt"

/*
	https://leetcode.com/problems/reverse-vowels-of-a-string/
	Given a string s, reverse only all the vowels in the string and return it.
	eg : "hello", Output: "holle"
*/

func reverseVowels(s string) string {
	ch := []byte(s)
	i, j := 0, len(s)-1
	for i <= j {
		iIsVowel := isVowel(ch[i])
		jIsVowel := isVowel(ch[j])
		if iIsVowel && jIsVowel {
			ch[i], ch[j] = ch[j], ch[i]
			i++
			j--
		} else if !iIsVowel {
			i++
		} else if !jIsVowel {
			j--
		}
	}
	return string(ch)
}

func isVowel(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
		return true
	}
	return false
}

func main() {
	s := "leetcode"
	fmt.Println(reverseVowels(s))
}
