package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Leetcode #226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// swap left and right
	var tmp = root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
