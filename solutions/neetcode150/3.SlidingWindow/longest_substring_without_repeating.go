package main

// leetcode #3
func lengthOfLongestSubstring(s string) int {
	l, r, maxLen := 0, 0, 0
	strMap := make(map[byte]bool)

	for r < len(s) {
		for strMap[s[r]] {
			strMap[s[l]] = false
			l++
		}

		strMap[s[r]] = true
		maxLen = Max(maxLen, r - l + 1)
		r++
	}

	return maxLen
}
