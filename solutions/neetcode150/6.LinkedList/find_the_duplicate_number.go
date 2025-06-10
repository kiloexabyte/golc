package main

// Leetcode #287
func findDuplicate(nums []int) int {
    var slow = 0
	var fast = 0

	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow = 0

	for true {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			return slow
		}
	}

	return -1
}