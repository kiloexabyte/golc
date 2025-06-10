package intervals

import "sort"

// Leetcode #56
func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    // Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    output := [][]int{intervals[0]}

    for i := 1; i < len(intervals); i++ {
        lastEnd := output[len(output)-1][1]
        start := intervals[i][0]
        end := intervals[i][1]

        if start <= lastEnd {
            if end > lastEnd {
                output[len(output)-1][1] = end
            }
        } else {
            output = append(output, []int{start, end})
        }
    }

    return output
}
