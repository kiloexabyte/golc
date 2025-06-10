package main

// Leetcode #347
func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{} // map of value and occurence
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		countMap[n]++
	}

	for key, value := range countMap {
		freq[value] = append(freq[value], key)
	}

	res := []int{}

	for i := len(freq) - 1; i > -1; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
