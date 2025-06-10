package graphs

// Leetcode #684
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	rank := make([]int, n+1)

	// Initialize parent and rank
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(n int) int {
		p := parent[n]
		for p != parent[p] {
			parent[p] = parent[parent[p]] // Path compression
			p = parent[p]
		}
		return p
	}

	union := func(n1, n2 int) bool {
		p1 := find(n1)
		p2 := find(n2)

		if p1 == p2 {
			return false // cycle detected
		}

		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return true
	}

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		if !union(n1, n2) {
			return edge
		}
	}

	return []int{}
}
