package dp2d

// Leetcode #494
func findTargetSumWays(nums []int, target int) int {
	result := 0

	var dfs func(index int, total int)
	dfs = func(index int, total int) {
		if index == len(nums) {
			if total == target {
				result++
			}
			return
		}

		dfs(index+1, total+nums[index])
		dfs(index+1, total-nums[index])
	}

	dfs(0, 0)
	return result
}

