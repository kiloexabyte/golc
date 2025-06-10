package dp

// letcode #647
func countSubstrings(s string) int {
    res := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// Odd length palindromes
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}

		// Even length palindromes
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	return res
}