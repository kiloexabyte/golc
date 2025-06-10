package greedy

// Leetcode #53
func maxSubArray(nums []int) int {
    maxSub := nums[0]
	curSum := 0

	for _, n := range nums {
		if curSum < 0 {
			curSum = 0
		}

		curSum += n
		maxSub = max(maxSub, curSum)
	}

	return maxSub

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}