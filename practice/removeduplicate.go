package practice

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev, size, i := nums[0], 1, 1

	for j := 1; j < len(nums); j++ {
		if prev != nums[j] {
			nums[i] = nums[j]
			i++
			size++
			prev = nums[j]
		}
	}
	return size
}
