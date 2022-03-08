package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Accepted, Runtime: 4 ms, Memory Usage: 4.3 MB
// Runtime: 4 ms, faster than 98.14% of Go online submissions for Linked List Cycle.
// Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Linked List Cycle.
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
