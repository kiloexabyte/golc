package main

import (
	"container/heap"
)

// Leetcode #703
// MinHeap is a type that implements heap.Interface for a min-heap of integers.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KthLargest represents the Kth largest element tracker.
type KthLargest struct {
	heap *MinHeap
	k    int
}

// Constructor initializes the KthLargest with a given k and nums array.
func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)
	kthLargest := KthLargest{
		heap: h,
		k:    k,
	}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

// Add adds a new value and returns the Kth largest element.
func (kl *KthLargest) Add(val int) int {
	heap.Push(kl.heap, val)
	if kl.heap.Len() > kl.k {
		heap.Pop(kl.heap)
	}
	return (*kl.heap)[0] // The root of the heap is the Kth largest element.
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */