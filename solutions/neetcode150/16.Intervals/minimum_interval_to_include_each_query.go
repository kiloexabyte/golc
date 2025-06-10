package intervals

import (
    "container/heap"
    "sort"
)

// Define a type for the heap elements: [intervalSize, intervalEnd]
type IntervalHeap [][]int

func (h IntervalHeap) Len() int { return len(h) }
func (h IntervalHeap) Less(i, j int) bool {
    // Compare by interval size (first element)
    return h[i][0] < h[j][0]
}
func (h IntervalHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntervalHeap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *IntervalHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Leetcode #1851
func minInterval(intervals [][]int, queries []int) []int {
    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // Copy and sort queries
    sortedQueries := make([]int, len(queries))
    copy(sortedQueries, queries)
    sort.Ints(sortedQueries)

    resMap := make(map[int]int)
    h := &IntervalHeap{}
    heap.Init(h)

    i := 0
    for _, q := range sortedQueries {
        // Add intervals that start before or at query q
        for i < len(intervals) && intervals[i][0] <= q {
            length := intervals[i][1] - intervals[i][0] + 1
            heap.Push(h, []int{length, intervals[i][1]})
            i++
        }

        // Remove intervals that end before query q
        for h.Len() > 0 && (*h)[0][1] < q {
            heap.Pop(h)
        }

        if h.Len() > 0 {
            resMap[q] = (*h)[0][0]
        } else {
            resMap[q] = -1
        }
    }

    // Build result for queries in original order
    result := make([]int, len(queries))
    for idx, q := range queries {
        result[idx] = resMap[q]
    }
    return result
}
