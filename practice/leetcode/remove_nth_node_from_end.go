package practice

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr, disFromStart, size := head, 0, 0

	for curr != nil {
		size++
		curr = curr.Next
	}

	disFromStart, curr = size-n, head
	prev := new(ListNode)
	for k := 0; k < disFromStart; k++ {
		prev = curr
		curr = curr.Next
	}

	if disFromStart == 0 {
		return head.Next
	}

	prev.Next = curr.Next

	return head
}
