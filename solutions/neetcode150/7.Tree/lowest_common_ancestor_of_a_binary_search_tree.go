package main

// Leetcode #235
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var curr = root

	for curr != nil {
		if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else {
			return curr
		}
	}

	return curr
}