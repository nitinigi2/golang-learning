// https://www.hackerrank.com/challenges/string-construction/problem
package main

func stringConstruction(s string) int32 {
	var count int32 = 0
	var arr [26]int

	for _, v := range s {
		arr[v-'a']++
	}

	for i := range arr {
		if arr[i] != 0 {
			count++
		}
	}
	return count
}
