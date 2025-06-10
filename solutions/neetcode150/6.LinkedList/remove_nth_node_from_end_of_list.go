package main

// Leetcode #19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	// Create a dummy node pointing to the head
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// Move `first` pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		if first == nil {
			return head // Return original head if n is out of bounds
		}
		first = first.Next
	}

	// Move both pointers until `first` reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Check if `second.Next` is valid before removing the node
	if second.Next != nil {
		second.Next = second.Next.Next
	}

	return dummy.Next // Return the new head of the list
}