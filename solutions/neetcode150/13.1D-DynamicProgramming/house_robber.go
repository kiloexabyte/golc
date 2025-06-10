package dp

// leetcode #198
func rob(nums []int) int {
    r1, r2 := 0,0

	for _, n := range nums {
		tmp := max(r1+n, r2)
		r1 = r2
		r2 = tmp
	}

	return r2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}