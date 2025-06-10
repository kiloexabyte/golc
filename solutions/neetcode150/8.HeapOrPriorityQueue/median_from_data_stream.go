package main

import (
	"container/heap"
)


// Leetcode #295
// MaxHeap stores the smaller half (as a max heap)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Peek() int {
	return h[0]
}

// MinHeap stores the larger half (as a min heap)
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MinHeap) Peek() int {
	return h[0]
}

// MedianFinder stores two heaps to track the median
type MedianFinder struct {
	small *MaxHeap
	large *MinHeap
}

func Constructor() MedianFinder {
	s := &MaxHeap{}
	l := &MinHeap{}
	heap.Init(s)
	heap.Init(l)
	return MedianFinder{
		small: s,
		large: l,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.small, num)

	// Move top of small to large if necessary
	if mf.small.Len() > 0 && mf.large.Len() > 0 && mf.small.Peek() > mf.large.Peek() {
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	}

	// Balance the sizes
	if mf.small.Len() > mf.large.Len()+1 {
		val := heap.Pop(mf.small).(int)
		heap.Push(mf.large, val)
	} else if mf.large.Len() > mf.small.Len() {
		val := heap.Pop(mf.large).(int)
		heap.Push(mf.small, val)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.small.Len() > mf.large.Len() {
		return float64(mf.small.Peek())
	} else if mf.large.Len() > mf.small.Len() {
		return float64(mf.large.Peek())
	} else {
		return (float64(mf.small.Peek()) + float64(mf.large.Peek())) / 2.0
	}
}