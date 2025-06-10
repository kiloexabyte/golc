package main

// Leetcode #74
func searchMatrix(matrix [][]int, target int) bool {
	var nums []int

	for _, row := range matrix {
		nums = append(nums, row...)
	}

	var min = 0
	var max = len(nums) - 1

	for min <= max {
		var mid = min + (max-min)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return false
}
