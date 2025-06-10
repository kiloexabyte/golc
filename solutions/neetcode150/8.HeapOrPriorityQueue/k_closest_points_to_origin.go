package main

import (
	"container/heap"
	"math"
)

// Leetcode	973
// PointHeap is a max-heap of points based on their distance from the origin.
type PointHeap struct {
	points   [][]int
	distances []float64
}

func (h *PointHeap) Len() int {
	return len(h.points)
}

func (h *PointHeap) Less(i, j int) bool {
	// Reverse order to make it a max-heap
	return h.distances[i] > h.distances[j]
}

func (h *PointHeap) Swap(i, j int) {
	h.points[i], h.points[j] = h.points[j], h.points[i]
	h.distances[i], h.distances[j] = h.distances[j], h.distances[i]
}

func (h *PointHeap) Push(x interface{}) {
	point := x.([]int)
	h.points = append(h.points, point)
	h.distances = append(h.distances, calculateDistance(point))
}

func (h *PointHeap) Pop() interface{} {
	lastIdx := len(h.points) - 1
	point := h.points[lastIdx]
	h.points = h.points[:lastIdx]
	h.distances = h.distances[:lastIdx]
	return point
}

// calculateDistance computes the Euclidean distance of a point from the origin.
func calculateDistance(point []int) float64 {
	xPart := float64(point[0] * point[0])
	yPart := float64(point[1] * point[1])
	return math.Sqrt(xPart + yPart)
}

// KClosest returns the k closest points to the origin.
func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{
		points:   make([][]int, 0, k),
		distances: make([]float64, 0, k),
	}
	heap.Init(h)

	for _, point := range points {
		heap.Push(h, point)
		if h.Len() > k {
			heap.Pop(h) // Remove the farthest point if the heap exceeds size k
		}
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).([]int))
	}

	return result
}
