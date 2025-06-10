package main

// Leetcode #143
func reorderList(head *ListNode)  {
    var pSlow = head
	var pFast = head.Next

	// finding the second half
	for pFast != nil && pFast.Next != nil {
		pSlow = pSlow.Next
		pFast = pFast.Next.Next
	}

	var second = pSlow.Next
	pSlow.Next = nil
	var prev = pSlow.Next

	// reverse second half
	for second != nil {
		var tmp = second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// merge two halves
	second = prev
	var first = head

	for second != nil {
		var tmp1 = first.Next
		var tmp2 = second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}