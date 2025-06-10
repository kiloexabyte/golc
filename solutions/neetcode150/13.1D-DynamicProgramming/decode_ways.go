package dp

// Leetcode #91
func numDecodings(s string) int {
	n := len(s)
	cache := make(map[int]int)
	cache[n] = 1 // Base case: empty string has one valid decoding

	var dfs func(int) int
	dfs = func(i int) int {
		if val, ok := cache[i]; ok {
			return val
		}
		if s[i] == '0' {
			return 0
		}

		res := dfs(i + 1)

		if i+1 < n && (s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6')) {
			res += dfs(i + 2)
		}

		cache[i] = res
		return res
	}

	return dfs(0)
}
