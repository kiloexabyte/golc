package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Leetcode #25
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Next: head}
	groupPrev := dummyNode

	for {
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next

		// Reverse the group
		var prev *ListNode = groupNext
		curr := groupPrev.Next

		for curr != groupNext {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}

		temp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = temp
	}

	return dummyNode.Next
}

func getKthNode(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}
