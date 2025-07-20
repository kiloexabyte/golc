package dp2d

// Leetcode #10
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make(map[[2]int]bool)

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        key := [2]int{i, j}
        if val, ok := dp[key]; ok {
            return val
        }
        if j == n {
            return i == m
        }

        firstMatch := i < m && (s[i] == p[j] || p[j] == '.')
        var ans bool

        if j+1 < n && p[j+1] == '*' {
            // 0 occurrence or 1+ occurrence
            ans = dfs(i, j+2) || (firstMatch && dfs(i+1, j))
        } else {
            ans = firstMatch && dfs(i+1, j+1)
        }

        dp[key] = ans
        return ans
    }

    return dfs(0, 0)
}