package main

func longestConsecutive(nums []int) int {
    numSet := make(map[int]struct{})
    for _, num := range nums {
        numSet[num] = struct{}{}
    }

    maxLen := 0

    for num := range numSet {
        // Only start counting if num is the beginning of a sequence
        if _, exists := numSet[num-1]; !exists {
            currentNum := num
            currentLen := 1

            // Use a while-like loop to check the condition
            for {
                if _, exists := numSet[currentNum+1]; exists {
                    currentNum++
                    currentLen++
                } else {
                    break
                }
            }

            if currentLen > maxLen {
                maxLen = currentLen
            }
        }
    }

    return maxLen
}