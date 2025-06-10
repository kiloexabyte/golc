package main

// Leetcode #42
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	l := 0
	r := len(height) - 1
	leftMax := height[l]
	rightMax := height[r]
	res := 0

	for l < r {
		if leftMax < rightMax { 
			l++
			leftMax = Max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = Max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}

	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}