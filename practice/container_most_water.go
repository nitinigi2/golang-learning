package practice

import "math"

/*
https://leetcode.com/problems/container-with-most-water/
*/
func maxArea(height []int) int {
	var i, maxWater, min int
	j := len(height) - 1
	for i <= j {
		min = int(math.Min(float64(height[i]), float64(height[j])))
		maxWater = int(math.Max(float64(maxWater), float64(min*(j-i))))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxWater
}
