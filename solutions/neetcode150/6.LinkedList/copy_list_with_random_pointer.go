package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Leetcode #138
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var nodeMap = map[*Node]*Node{}
	var current = head

	for current != nil {
		var copy = &Node{Val: current.Val}
		nodeMap[current] = copy

		current = current.Next
	}

	current = head

	for current != nil {
		var copyNode = nodeMap[current]

		if current.Next != nil {
			var nextNode = nodeMap[current.Next]
			copyNode.Next = nextNode
		} else {
			copyNode.Next = nil
		}

		if current.Random != nil {
			var randomNode = nodeMap[current.Random]
			copyNode.Random = randomNode
		} else {
			copyNode.Random = nil
		}

		current = current.Next
	}

	return nodeMap[head]
}
