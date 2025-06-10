package main

import (
	"container/heap"
)

// lc #1046
// MaxHeap is a type that implements heap.Interface for a max-heap of integers.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// LastStoneWeight calculates the weight of the last stone.
func LastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	// Add all stones to the heap (negate values to simulate max-heap)
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	// Smash stones until one or none remain
	for h.Len() > 1 {
		y := heap.Pop(h).(int) // Heavier stone
		x := heap.Pop(h).(int) // Lighter stone

		if y > x {
			heap.Push(h, y-x)
		}
	}

	// If no stones remain, return 0. Otherwise, return the weight of the last stone.
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}
