package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

// Leetcode #1448
func goodNodes(root *TreeNode) int {
	var goodNodeCount int // Local variable

	var dfs func(node *TreeNode, max int)
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}

		newMax := max
		if node.Val >= max {
			goodNodeCount++
			newMax = node.Val
		}

		dfs(node.Left, newMax)
		dfs(node.Right, newMax)
	}

	dfs(root, math.MinInt)
	return goodNodeCount
}