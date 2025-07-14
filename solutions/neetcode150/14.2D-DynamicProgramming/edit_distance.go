package dp2d

// Leetcode #72
func minDistance(word1 string, word2 string) int {
	n1 := len(word1)
	n2 := len(word2)
    cache := make([][]int, n1 + 1)
	
	for i := range cache {
		cache[i] = make([]int, n2 + 1)
	}

	// base case
	for j := range n2 + 1 {
		cache[n1][j] = n2 - j
	}

	for i := range n1 + 1 {
		cache[i][n2] = n1 - i
	}

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				cache[i][j] = cache[i + 1][j + 1]
			} else {
				cache[i][j] = 1 + min(cache[i + 1][j], cache[i][j + 1], cache[i + 1][j + 1])
			}
		}
	}

	return cache[0][0]
}

func min(a, b, c int) int {
    if a < b && a < c {
        return a
    }
    if b < c {
        return b
    }
    return c
}