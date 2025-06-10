package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Leetcode #105
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Root is the first element of preorder traversal
	root := &TreeNode{Val: preorder[0]}

	// Find the index of the root in inorder traversal
	mid := 0
	for i, val := range inorder {
		if val == preorder[0] {
			mid = i
			break
		}
	}

	// Recursively build the left and right subtrees
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
