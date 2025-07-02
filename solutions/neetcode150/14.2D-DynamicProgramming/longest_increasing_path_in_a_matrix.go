package dp2d

// Leetcode #329
func longestIncreasingPath(matrix [][]int) int {
    rows :=  len(matrix)
	cols := len(matrix[0])
	maxVal := 0

	dp := map[[2]int]int{}

	var dfs func(r, c, prevVal int) int
	dfs = func(r,c, prevVal int) int {
		if r < 0 || r == rows || c < 0 || c == cols || matrix[r][c] <= prevVal {
			return 0
		}

		key := [2]int{r, c}
		_, exist := dp[key]
		if exist {
			return dp[key]
		}

		res := 1
		res = max(res, dfs(r + 1, c, matrix[r][c]) + 1)
		res = max(res, dfs(r - 1, c, matrix[r][c]) + 1)
		res = max(res, dfs(r, c + 1, matrix[r][c]) + 1)
		res = max(res, dfs(r, c - 1, matrix[r][c]) + 1)

		dp[key] = res
		maxVal = max(maxVal, res)
		return res
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(r,c, -1)
		}
	}

	return maxVal
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}