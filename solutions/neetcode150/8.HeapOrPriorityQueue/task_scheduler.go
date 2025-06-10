package main

// Leetcode #621
func leastInterval(tasks []byte, n int) int {
	// Step 1: Calculate frequencies of each task
	taskCounts := make([]int, 26)
	for _, task := range tasks {
		taskCounts[task-'A']++
	}

	// Step 2: Find the maximum frequency
	maxFreq := 0
	for _, count := range taskCounts {
		if count > maxFreq {
			maxFreq = count
		}
	}

	// Step 3: Count the number of tasks with maximum frequency
	maxFreqCount := 0
	for _, count := range taskCounts {
		if count == maxFreq {
			maxFreqCount++
		}
	}

	// Step 4: Calculate the minimum intervals needed
	partCount := maxFreq - 1
	partLength := n - (maxFreqCount - 1)
	emptySlots := partCount * partLength
	availableTasks := len(tasks) - maxFreq*maxFreqCount
	idles := max(0, emptySlots-availableTasks)

	return len(tasks) + idles
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}