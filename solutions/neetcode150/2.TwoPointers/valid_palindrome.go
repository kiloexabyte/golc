package main

import (
	"strings"
	"unicode"
)

func cleanString(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}

// leetcode #125
func isPalindrome(s string) bool {
	s = cleanString(s)
	p1 := 0
	p2 := len(s) - 1


	for p1 < p2 {
		if s[p1] != s[p2] {
			return false
		}

		p1++
		p2--
	}

	return true
}
