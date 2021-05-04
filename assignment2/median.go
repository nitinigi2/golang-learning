// https://www.hackerrank.com/challenges/find-the-median/problem
package main

import "sort"

func findMedian(arr []int32) int32 {
	n := len(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if n%2 == 0 {
		return arr[n/2]
	}
	return (arr[n/2] + arr[n/2+1]) / 2
}
