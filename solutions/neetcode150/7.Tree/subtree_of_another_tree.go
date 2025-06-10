package main

// Leetcode #572
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
		return true
	}

	if root == nil { 
		return false
	}

	if sameTree(root,subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return sameTree(p.Right, q.Right) && sameTree(p.Left, q.Left)
}