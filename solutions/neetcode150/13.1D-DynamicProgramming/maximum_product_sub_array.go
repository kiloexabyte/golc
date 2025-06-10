package dp

// Leetcode #152
func maxProduct(nums []int) int {
    res := maxnum(nums)

	curMin := 1
	curMax := curMin

	for _, n := range nums {
		if n == 0 {
			curMin = 1
			curMax = 1
			continue
		}

		temp := curMax * n
		curMax = Max(Max(n*curMax,n*curMin), n)
		curMin = Min(Min(temp, n*curMin), n)

		res = Max(res, curMax)
	}

	return res

}

func maxnum(nums []int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {	
	if a < b {
		return a
	}
	return b
}
