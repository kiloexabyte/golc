package main

var (
	r [][]int
	t int
	cs []int
)


// Leetcode #39
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var dfs func(i int, curr []int, total int)

	dfs = func(i int, curr []int, total int) {
		if total == target {
			// Make a copy of curr before adding
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			res = append(res, tmp)
			return
		}

		if i == len(candidates) || total > target {
			return
		}

		// include current i value
		curr = append(curr, candidates[i])
		dfs(i, curr, total+candidates[i])

		// backtrack: remove last added element
		curr = curr[:len(curr)-1]
		dfs(i+1, curr, total)
	}

	dfs(0, []int{}, 0)
	return res
}