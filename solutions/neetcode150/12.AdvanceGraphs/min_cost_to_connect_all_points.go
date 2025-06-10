package advancegraphs

import (
	"container/heap"
)

// lc #1584
func minCostConnectPoints(points [][]int) int {
	N := len(points)

	// Build adjacency list with cost to connect each point
	adjList := make([][][2]int, N)
	for i := 0; i < N; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < N; j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := abs(x1 - x2) + abs(y1 - y2)

			adjList[i] = append(adjList[i], [2]int{j, dist})
			adjList[j] = append(adjList[j], [2]int{i, dist})
		}
	}

	visited := make([]bool, N)
	res := 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Value: 0, Priority: 0}) // start from node 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		node := item.Value
		cost := item.Priority

		if visited[node] {
			continue
		}

		visited[node] = true
		res += cost

		for _, pair := range adjList[node] {
			nei, neiCost := pair[0], pair[1]
			if !visited[nei] {
				heap.Push(pq, &Item{Value: nei, Priority: neiCost})
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// --- MinHeap code below ---

type Item struct {
	Value    int // node index
	Priority int // cost to connect
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority // MinHeap
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
