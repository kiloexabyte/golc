package main

// Leetcode #11
func maxArea(height []int) int {
    l := 0
	r := len(height) - 1
	areaMax := 0

	for l < r {
		area := (r-l) * Min(height[r], height[l])
		areaMax = Max(area, areaMax)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return areaMax
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}