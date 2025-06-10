package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}
