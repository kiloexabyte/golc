package main

// leetcode #242
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if charMap[t[i]] <= 0 {
			return false
		}
		charMap[t[i]]--
	}

	return true
}
