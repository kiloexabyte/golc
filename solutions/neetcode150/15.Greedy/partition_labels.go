package greedy

// Leetcode #763
func partitionLabels(s string) []int {
    hm := map[byte]int{} // last character's index map
	n := len(s)
	for i := 0; i < n; i++ {
		hm[s[i]] = i
	}

	size := 0
	end := 0
	output := []int{};
	for i := 0; i < n; i++ {
		size++
		end = max(end, hm[s[i]])
		if i == end {
			output = append(output, size)
			end = 0
			size = 0
		}
	}

	return output
}

func max(a,b int) int {
	if a > b {
		return a	
	}
	return b
}