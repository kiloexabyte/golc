package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// Leetcode #199
func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{}
	levels := [][]int{}
	res := []int{}

	queue = append(queue, root)

	for len(queue) > 0 {
		var qLen = len(queue)
		var level = []int{ -1 }

		for i := 0; i < qLen; i++ {
			var node = queue[0]
			queue = queue[1:]

			if node != nil {
				level[0] = node.Val
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		if level[0] != -1 {
			levels = append(levels, level)
		}
	}

	for _, l := range levels {
		res = append(res, l[len(l)-1])
	}

	return res

}
