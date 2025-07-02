package dp2d


// Leetcode #115
func numDistinct(s string, t string) int {
    cache := map[[2]int]int{}

	var dfs func(i,j int) int
	dfs = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		key := [2]int{i, j}
		val, exist := cache[key]
		if exist {
			return val
		}

		if s[i] == t[j] {
			cache[key] = dfs(i + 1, j + 1) + dfs(i + 1, j)
		} else {
			cache[key] = dfs(i + i, j)
		}
		return cache[key]
	} 

	return dfs(0,0)
}