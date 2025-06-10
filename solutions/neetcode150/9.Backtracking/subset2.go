package main

import "sort"

var rs [][]int

// Leetcode #90
func subsetsWithDup(nums []int) [][]int {
	rs = [][]int{}
	sort.Ints(nums)
	backtrack(0, []int{}, nums)
	return rs
}

func backtrack(i int, subset []int, nums []int) {
	if i == len(nums) {
		tmp := make([]int, len(subset))
		copy(tmp, subset)
		rs = append(rs, tmp)
		return
	}

	// Include nums[i]
	subset = append(subset, nums[i])
	backtrack(i+1, subset, nums)
	subset = subset[:len(subset)-1] // backtrack

	// Skip duplicates
	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}

	// Exclude nums[i]
	backtrack(i+1, subset, nums)
}
