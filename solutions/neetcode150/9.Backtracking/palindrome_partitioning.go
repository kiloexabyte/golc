package main


// Leetcode #131
func partition(s string) [][]string {
	sLen := len(s)

	rslt := [][]string{}
	part := []string{}

	var dfs func(i int)
    dfs = func (i int) {
		if i == sLen {
			tmp := make([]string, len(part))
			copy(tmp, part)
			rslt = append(rslt, tmp)
			return
		}

		for j := i; j < sLen; j++ {
			if isPalindrome(s, i, j) {
				part = append(part, s[i : j+1])
				dfs(j + 1)
				part = part[:len(part)-1]
			}
		}
	}

	dfs(0)
	return rslt

}

func isPalindrome(str string, i,j int) bool {
	l := i
	r := j
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}