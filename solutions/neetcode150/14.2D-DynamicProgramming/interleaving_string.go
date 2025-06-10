package dp2d

// Leetcode #97
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make(map[[2]int]bool)

	// Initialize all positions to false
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			dp[[2]int{i, j}] = false
		}
	}

	dp[[2]int{len(s1), len(s2)}] = true

	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {
			if i < len(s1) && s1[i] == s3[i+j] && dp[[2]int{i + 1, j}] {
				dp[[2]int{i, j}] = true
			}
			if j < len(s2) && s2[j] == s3[i+j] && dp[[2]int{i, j + 1}] {
				dp[[2]int{i, j}] = true
			}
		}
	}

	return dp[[2]int{0, 0}]
}
