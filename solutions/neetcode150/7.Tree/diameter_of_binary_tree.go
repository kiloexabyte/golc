package main

// Leetcode 543
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		
		left := dfs(node.Left)
		right := dfs(node.Right)
		
		// Update the diameter
		res = max(res, 2+left+right)

		// Return the height of the current node
		return 1 + max(left, right)
	}

	dfs(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}