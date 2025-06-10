package advancegraphs

import (
	"container/heap"
)

// Leetcode #743
func networkDelayTime(times [][]int, n int, k int) int {
	// Build adjacency list: node -> list of (neighbor, weight)
	graph := make(map[int][][2]int)
	for _, t := range times {
		src, dst, wt := t[0], t[1], t[2]
		graph[src] = append(graph[src], [2]int{dst, wt})
	}

	// Min-heap: (time, node)
	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, [2]int{0, k})

	visited := make(map[int]bool)
	dist := make(map[int]int)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).([2]int)
		time, node := curr[0], curr[1]

		if visited[node] {
			continue
		}
		visited[node] = true
		dist[node] = time

		for _, edge := range graph[node] {
			neighbor, weight := edge[0], edge[1]
			if !visited[neighbor] {
				heap.Push(pq, [2]int{time + weight, neighbor})
			}
		}
	}

	if len(dist) != n {
		return -1
	}

	maxTime := 0
	for _, t := range dist {
		if t > maxTime {
			maxTime = t
		}
	}
	return maxTime
}

// ------------------ MinHeap implementation ------------------

type MinHeap [][2]int // [time, node]

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] } // min-heap based on time
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
