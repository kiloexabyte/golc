package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Leetcode #230
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	var n = 0
	var curr = root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop the stack
		n++

		if k == n {
			return curr.Val
		}

		curr = curr.Right
	}

	return -1
}
