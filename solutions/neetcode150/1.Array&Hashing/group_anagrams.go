package main

import (
	"sort"
	"strings"
)

// leetcode #49
func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)

	for _, str := range strs {
		// Convert string to a slice of runes for sorting
		charArr := strings.Split(str, "")
		sort.Strings(charArr)
		key := strings.Join(charArr, "")

		// Append the string to the correct group
		res[key] = append(res[key], str)
	}

	// Convert map values to a slice of slices
	result := make([][]string, 0, len(res))
	for _, group := range res {
		result = append(result, group)
	}

	return result
}
