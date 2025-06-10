package main

type Pair struct {
	x int
	y int
}

var (
	rows int
	cols int
	path map[Pair]bool
)

// Leetcode #79
func exist(board [][]byte, word string) bool {
	rows = len(board)
	cols = len(board[0])
	path = make(map[Pair]bool)

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		p := Pair{r, c}
		if r < 0 || c < 0 || r >= rows || c >= cols || word[i] != board[r][c] || path[p] {
			return false
		}

		path[p] = true

		res := dfs(r+1, c, i+1) ||
			dfs(r-1, c, i+1) ||
			dfs(r, c+1, i+1) ||
			dfs(r, c-1, i+1)

		delete(path, p) // backtrack

		return res
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
