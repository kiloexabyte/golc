package main

// Leetcode #33
func search(nums []int, target int) int {
    var l = 0
	var r = len(nums) - 1

	for l <= r {
		var m = (l + r) / 2
		if target == nums[m] {
			return m
		}

		if nums[l] <= nums[m] {
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}