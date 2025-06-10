package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Leetcode #297
type Codec struct {
	resultList []string
	index      int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.resultList = []string{}
	this.dfsSerialize(root)
	return strings.Join(this.resultList, ",")
}

func (this *Codec) dfsSerialize(node *TreeNode) {
	if node == nil {
		this.resultList = append(this.resultList, "N")
		return
	}

	this.resultList = append(this.resultList, strconv.Itoa(node.Val))
	this.dfsSerialize(node.Left)
	this.dfsSerialize(node.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	this.index = 0
	return this.dfsDeserialize(vals)
}

func (this *Codec) dfsDeserialize(vals []string) *TreeNode {
	if vals[this.index] == "N" {
		this.index++
		return nil
	}

	val, _ := strconv.Atoi(vals[this.index])
	node := &TreeNode{Val: val}
	this.index++
	node.Left = this.dfsDeserialize(vals)
	node.Right = this.dfsDeserialize(vals)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
