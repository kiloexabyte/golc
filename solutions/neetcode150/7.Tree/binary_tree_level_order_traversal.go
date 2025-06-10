package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // Leetcode #102
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}

    if root == nil {
        return res
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        qLen := len(queue)
        level := []int{}

        for i := 0; i < qLen; i++ {
            node := queue[0]
            queue = queue[1:] // Dequeue the first element

            if node != nil {
                level = append(level, node.Val)
                if node.Left != nil {
                    queue = append(queue, node.Left)
                }
                if node.Right != nil {
                    queue = append(queue, node.Right)
                }
            }
        }

        if len(level) > 0 {
            res = append(res, level)
        }
    }

    return res
}
