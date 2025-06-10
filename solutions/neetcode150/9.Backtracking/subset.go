package main

// Leetcode #78
var (
	res    [][]int
	subset []int
	numss   []int
)

func subsets(nums []int) [][]int {
	res = [][]int{}
	subset = []int{}
	numss = nums
	dfs(0)
	return res
}

func dfs(i int) {
	if i == len(numss) {
		// Make a copy of the current subset
		copySubset := make([]int, len(subset))
		copy(copySubset, subset)
		res = append(res, copySubset)
		return
	}

	// Include nums[i]
	subset = append(subset, numss[i])
	dfs(i + 1)

	// Exclude nums[i]
	subset = subset[:len(subset)-1]
	dfs(i + 1)
}