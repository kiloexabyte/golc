package main

// leetcode #424
func characterReplacement(s string, k int) int {
	l, r, maxf, res := 0, 0, 0, 0
	seenMap := map[byte]int{}

	for r < len(s) {
		// figure out max freq letter
		seenMap[s[r]]++
		maxf = Max(maxf, seenMap[s[r]])

		// move L up until subtitution necessary is equal to k
		for r - l + 1 - maxf > k {
			seenMap[s[l]]--
			l++
		} 
        
		// calculate new max window length
		res = Max(res, r - l + 1)
		r++
	}

	return res
}