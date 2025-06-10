package greedy

// Leetcode #45
func jump(nums []int) int {
    res, left, right := 0, 0, 0

	for right < len(nums) - 1 {
		farthest := 0
		for i := left; i <= right; i++ {
			farthest = max(farthest, i + nums[i])
		}
		left = right + 1
		right = farthest
		res++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}