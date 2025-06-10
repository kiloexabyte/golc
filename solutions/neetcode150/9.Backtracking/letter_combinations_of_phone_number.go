package main

var dict = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// Leetcode #17
func letterCombinations(digits string) []string {
    res := []string{}
    var backtrack func (i int, currStr string, digits string)
    backtrack = func (i int, currStr string, digits string) {
        if len(currStr) == len(digits) {
            res = append(res, currStr)
            return
        }

        // Recursively append letters corresponding to the digit at the current position
        for _, c := range dict[rune(digits[i])] {
            backtrack(i+1, currStr+string(c), digits)
        }
    }
    
	if len(digits) > 0 {
		backtrack(0, "", digits)
	}
	return res
}
