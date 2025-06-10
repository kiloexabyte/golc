package bitmanipulation

// Leetcode #268
func missingNumber(nums []int) int {
    r := 0
	n := len(nums) + 1
	s := make([]int, n)
	
	for i := 0; i < n; i++ {
		s[i] = i
	}

	for _, nu := range s {
		r = r ^ nu
	}

	for _, nu := range nums {
		r = r ^ nu
	}

	return r
}

