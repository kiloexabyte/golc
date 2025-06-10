package main

// Leetcode 239
func maxSlidingWindow(nums []int, k int) []int {
    var l = 0
	var r = 0

	var res = []int{}
	var q = []int{}

	for r < len(nums) {
		
		for len(q) > 0 && nums[q[len(q) - 1]] < nums[r] {
			q = removeAtIndex(q, len(q) - 1)
		}

		q = append(q, r)

		if l > q[0] {
			q = removeAtIndex(q, 0)
		}

		if r + 1 >= k {
			res = append(res,nums[q[0]])
			l++
		}

		r++
	}

	return res
}

func removeAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		// Return the original slice if the index is out of bounds
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}