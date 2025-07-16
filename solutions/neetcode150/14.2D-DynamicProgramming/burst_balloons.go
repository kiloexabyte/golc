package dp2d

// Leetcode #312
func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

	dp := make(map[[2]int]int)

	var dfs func(l, r int) int
    dfs = func(l, r int) int {
		if l > r {
			return 0
		}
		key := [2]int {l,r}
		if val, exists := dp[key]; exists {
			return val
		}

		dp[key] = 0
		for i := l; i < r + 1; i++ {
			coins := nums[l - 1] * nums[i] * nums[r + 1]
			coins += dfs(l, i - 1) + dfs(i + 1, r)
			dp[key] = max(dp[key], coins)
		}
		return dp[key]
    }

	return dfs(1, len(nums) - 2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}