package main

import "container/heap"

// Leetcode #215
// MaxHeap implements a max-heap for integers.
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

func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	// Push all elements into the max-heap
	for _, n := range nums {
		heap.Push(h, n)
	}

	// Extract the k largest elements
	for i := 1; i < k; i++ {
		heap.Pop(h)
	}

	// The kth largest element
	return heap.Pop(h).(int)
}
