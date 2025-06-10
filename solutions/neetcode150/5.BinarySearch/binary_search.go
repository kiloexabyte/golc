package main

// Leetcode 704
func search(nums []int, target int) int {
    min := 0
	max := len(nums) - 1

	for min <= max {
		mid := min + (max - min) / 2
		
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}