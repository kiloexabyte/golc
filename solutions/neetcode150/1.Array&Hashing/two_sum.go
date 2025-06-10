package main

// leetcode #1
func twoSum(nums []int, target int) []int {
    var prevMap = make(map[int]int) // create dictionary of value and index

	for i := 0; i < len(nums); i++ {
		var checkThisNumber = target - nums[i]

		if _, exists := prevMap[checkThisNumber]; exists {
			return []int{prevMap[checkThisNumber], i}
		} else {
			prevMap[nums[i]] = i
		}
	}

	return nums
}