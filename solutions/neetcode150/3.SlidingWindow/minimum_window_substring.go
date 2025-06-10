package main

import "math"

// Leetcode #76
func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	// Create frequency maps for `t` and the sliding window
	countT := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	have, need := 0, len(countT)
	res := [2]int{-1, -1} // Store the start and end indices of the result
	resLen := math.MaxInt32
	l := 0

	// Expand the right boundary of the window
	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++

		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		// Shrink the left boundary of the window when the condition is met
		for have == need {
			if (r - l + 1) < resLen {
				res = [2]int{l, r}
				resLen = r - l + 1
			}

			window[s[l]]--

			if countT[s[l]] > 0 && window[s[l]] < countT[s[l]] {
				have--
			}

			l++
		}
	}

	if resLen == math.MaxInt32 {
		return ""
	}

	return s[res[0] : res[1]+1]
}