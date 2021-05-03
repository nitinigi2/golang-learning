package practice

// https://leetcode.com/problems/palindrome-linked-list/

Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slice, curr := make([]int, 0), head

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	for curr != nil {
		last := slice[len(slice)-1]
		if curr.Val != last {
			return false
		}

		slice = slice[:len(slice)-1]
		curr = curr.Next
	}
	return true
}
