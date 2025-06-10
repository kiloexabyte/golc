package main

type KeyValue struct {
	Index  int
	Height int
}

// Leetcode 84
func largestRectangleArea(heights []int) int {
	var stack = []KeyValue{}
	var maxArea = 0

	for i := 0; i < len(heights); i++ {
		var start = i
		for len(stack) > 0 && stack[len(stack)-1].Height > heights[i] {
			var kvp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var index = kvp.Index
			var height = kvp.Height

			maxArea = Max(maxArea, height*(i-index))
			start = index
		}

		stack = append(stack, KeyValue{Index: start, Height: heights[i]})
	}

	for _, kvp := range stack {
		maxArea = Max(maxArea, kvp.Height*(len(heights)-kvp.Index))
	}

	return maxArea
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
