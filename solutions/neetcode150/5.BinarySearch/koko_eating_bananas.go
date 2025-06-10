package main

import "math"

// Leetcode #875
func minEatingSpeed(piles []int, h int) int {
	var l = 1
	var r = FindMax(piles)

	var res = r

	for l <= r {
		var k = (l + r) / 2
		var hours = 0

		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}

		if hours <= h {
			if hours < 0 {
				break
			}

			res = min(res, k)
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

// FindMax returns the maximum value in an int slice
func FindMax(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}

	maxVal := nums[0] // Assume the first element is the max initially

	for _, num := range nums[1:] {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}


func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}