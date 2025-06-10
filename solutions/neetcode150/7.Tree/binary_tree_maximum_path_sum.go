package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Leetcode #124
func maxPathSum(root *TreeNode) int {
	// Initialize res with the smallest possible value
	res := math.MinInt
	dfs2(root, &res)
	return res
}

func dfs2(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}

	// Recursively calculate the maximum path sum for left and right subtrees
	leftMax := max(dfs2(node.Left, res), 0) // Ignore negative paths
	rightMax := max(dfs2(node.Right, res), 0)

	// Update the result with the maximum path sum including the current node
	*res = max(*res, node.Val+leftMax+rightMax)

	// Return the maximum path sum where the current node is the endpoint
	return node.Val + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
