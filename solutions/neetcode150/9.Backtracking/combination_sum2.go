package main

import "sort"

// Leetcode #40
var (
	c []int
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	c = candidates

	var result [][]int
	var curr []int

	backtrack(curr, 0, target, &result)
	return result
}

func backtrack(curr []int, start, target int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
		return
	}

	if target < 0 {
		return
	}

	prev := -1
	for i := start; i < len(c); i++ {
		if c[i] == prev {
			continue
		}

		curr = append(curr, c[i])
		backtrack(curr, i+1, target-c[i], result)
		curr = curr[:len(curr)-1]

		prev = c[i]
	}
}