// https://www.hackerrank.com/challenges/arrays-ds/problem
package main

func reverseArray(arr []int32) []int32 {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
