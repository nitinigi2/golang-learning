// https://www.hackerrank.com/challenges/camelcase/problem
package main

import "unicode"

func camelcase(s string) int32 {
	var ans int32 = 1
	for _, v := range s {
		if unicode.IsUpper(v) {
			ans++
		}
	}
	return ans
}
