package main

// FindMax returns the maximum value in an int slice
func FindMax(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}

	maxVal := nums[0] // Assume the first element is the max initially

	for _, num := range nums[1:] {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}