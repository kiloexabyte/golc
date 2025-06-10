package main

// leetcode #121
func maxProfit(prices []int) int {
    l := 0
	r := 1
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = Max(maxProfit, profit)
		} else {
			l = r
		}
		r++
	}

	return maxProfit
}

func Max(a,b int) int {
	if a < b {
		return b
	}

	return a
}