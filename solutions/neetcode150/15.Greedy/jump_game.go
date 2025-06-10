package greedy

// Leetcode #55
func canJump(nums []int) bool {
	goalPost := len(nums) - 1

	for i := len(nums) - 1; i > -1; i-- {
		if i+nums[i] >= goalPost {
			goalPost = i
		}
	}

	return goalPost == 0
}