package main

// leetcode #217
func containsDuplicate(nums []int) bool {
	var existMap = map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := existMap[nums[i]]; ok {
			return true
		} else {
			existMap[nums[i]] = true
		}
	}

	return false;

}