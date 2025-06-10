package bitmanipulation

// Leetcode 136
func singleNumber(nums []int) int {
    var res int
	for _, n := range nums {
		res = res ^ n
	}

	return res
}