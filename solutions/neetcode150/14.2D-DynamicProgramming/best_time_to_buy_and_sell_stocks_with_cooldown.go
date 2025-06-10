package dp2d


// Leetcode #309
func maxProfit(prices []int) int {
	dp := make(map[[2]int]int)

	var dfs func(i int, buying bool) int
	dfs = func(i int, buying bool) int {
		if i >= len(prices) {
			return 0
		}

		key := [2]int{i, 0}
		if buying {
			key[1] = 1
		}

		if val, ok := dp[key]; ok {
			return val
		}

		coolDown := dfs(i+1, buying)
		if buying {
			buy := dfs(i+1, false) - prices[i]
			dp[key] = max(buy, coolDown)
		} else {
			sell := dfs(i+2, true) + prices[i]
			dp[key] = max(sell, coolDown)
		}
		return dp[key]
	}

	return dfs(0, true)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}