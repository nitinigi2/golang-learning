// https://www.hackerrank.com/challenges/array-left-rotation/problem
package main

func rotateLeft(d int32, arr []int32) []int32 {
	var i, n int32 = 0, int32(len(arr))
	//i+n-d
	brr := make([]int32, n)

	for ; i < n; i++ {
		brr[(i+n-d)%n] = arr[i]
	}
	return brr
}
