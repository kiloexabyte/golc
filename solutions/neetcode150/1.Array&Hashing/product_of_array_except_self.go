package main


// Leetcode #238
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	pre := 1
	post := 1

	result[0] = pre
	
	// calculate prefix product of index i
	for i := 1; i < numsLen; i++ {
		pre *= nums[i-1]
		result[i] = pre
	}

	for i := numsLen - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i] // calculate postfix product of index i
	}

	return result
}