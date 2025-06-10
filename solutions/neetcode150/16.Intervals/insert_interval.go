package intervals

// Leetcode #57
func insert(intervals [][]int, newInterval []int) [][]int {
    var res [][]int
    inserted := false

    for i := 0; i < len(intervals); i++ {
        // If newInterval ends before current interval starts, insert it here
        if newInterval[1] < intervals[i][0] {
            if !inserted {
                res = append(res, newInterval)
                inserted = true
            }
            res = append(res, intervals[i:]...)
            return res
        } else if newInterval[0] > intervals[i][1] {
            // Current interval ends before newInterval starts, add current interval
            res = append(res, intervals[i])
        } else {
            // Overlapping intervals, merge
            if intervals[i][0] < newInterval[0] {
                newInterval[0] = intervals[i][0]
            }
            if intervals[i][1] > newInterval[1] {
                newInterval[1] = intervals[i][1]
            }
        }
    }

    // If newInterval not inserted yet, add it at the end
    if !inserted {
        res = append(res, newInterval)
    }

    return res
}
