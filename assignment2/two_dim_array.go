// https://www.hackerrank.com/challenges/2d-array/problem
package main

import (
	"math"
)

func hourglassSum(arr [][]int32) int32 {
	max := int32(math.MinInt32)
	m, n := len(arr), len(arr[0])
	var currSum int32
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			currSum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if currSum > max {
				max = int32(currSum)
			}
		}
	}
	return max
}
