package intervals

import "sort"

// Leetcode #435
func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }

    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    res := 0
    prevEnd := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        end := intervals[i][1]

        if start >= prevEnd {
            prevEnd = end
        } else {
            res++
            if end < prevEnd {
                prevEnd = end
            }
        }
    }

    return res
}
