package main

// Leetcode #104
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}


func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}