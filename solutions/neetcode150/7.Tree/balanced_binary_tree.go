package main

import "math"

// Leetcode #110
func isBalanced(root *TreeNode) bool {
	balanced, _ := dfs(root)
	return balanced
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftHeight := dfs(root.Left)
	rightBalanced, rightHeight := dfs(root.Right)

	// Check if current node is balanced
	balanced := leftBalanced && rightBalanced && math.Abs(float64(leftHeight-rightHeight)) <= 1

	// Return whether it's balanced and the height
	return balanced, 1 + max(leftHeight, rightHeight)
}