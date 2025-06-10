package dp

// leetcode #139
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	cache := make(map[int]bool)
	for i := 0; i <= n; i++ {
		cache[i] = false
	}
	cache[n] = true

	for i := n - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if i+len(w) <= n && s[i:i+len(w)] == w {
				if cache[i+len(w)] {
					cache[i] = true
					break
				}
			}
		}
	}

	return cache[0]
}
