package dp

// leetcode #300
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	return maxSlice(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlice(arr []int) int {
	m := arr[0]
	for _, val := range arr[1:] {
		if val > m {
			m = val
		}
	}
	return m
}
