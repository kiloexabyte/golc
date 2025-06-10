package graphs

type Node struct {
	Val       int
	Neighbors []*Node
}


// Leetcode #133
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if cloned, ok := visited[n.Val]; ok {
			return cloned
		}

		clone := &Node{Val: n.Val}
		visited[n.Val] = clone

		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}
