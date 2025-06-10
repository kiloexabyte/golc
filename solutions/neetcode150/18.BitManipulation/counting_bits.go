package bitmanipulation

// Leetcode 338
func countBits(n int) []int {
	l := n + 1
	dp := make([]int, l)
	o := 1

	for i := 1; i < l; i++ {
		if o * 2 == i {
			o = i
		}
		dp[i] = 1 + dp[i - o]
	}

	return dp
}