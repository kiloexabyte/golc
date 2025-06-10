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

// Leetcode #98
func isValidBST(root *TreeNode) bool {
	negativeInfinity := math.Inf(-1) // Negative infinity
	positiveInfinity := math.Inf(1)
	return dfs1(root, negativeInfinity, positiveInfinity)
}

func dfs1(node *TreeNode, left float64, right float64) bool {
	if node == nil {
		return true
	}

	if !(float64(node.Val) < right && float64(node.Val) > left) {
		return false
	}

	var res = dfs1(node.Left, left, float64(node.Val)) && dfs1(node.Right, float64(node.Val), right)
	return res
}
