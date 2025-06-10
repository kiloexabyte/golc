package main

// Leetcode #153
func findMin(nums []int) int {
	var l = 0
	var r = len(nums) - 1
	var res = nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			res = min(res, nums[l])
			return res
		}

		var m = (l + r) / 2
		res = min(res, nums[m])

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
