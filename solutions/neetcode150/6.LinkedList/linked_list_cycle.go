package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // leetcode #141
func hasCycle(head *ListNode) bool {
	var pFast = head
	var pSlow = head

	for pFast.Next != nil {
		pFast = pFast.Next.Next
		pSlow = pSlow.Next
		if pSlow == pFast {
			return true
		}
	}

	return false
}