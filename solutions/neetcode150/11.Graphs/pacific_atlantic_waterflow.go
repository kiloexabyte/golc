package graphs

// leetcode #417
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)
	res := [][]int{}

	var dfs func(r, c int, visited map[[2]int]bool, prevHeight int)
	dfs = func(r, c int, visited map[[2]int]bool, prevHeight int) {
		if r < 0 || c < 0 || r >= rows || c >= cols {
			return
		}
		if heights[r][c] < prevHeight {
			return
		}
		key := [2]int{r, c}
		if visited[key] {
			return
		}

		visited[key] = true

		dfs(r+1, c, visited, heights[r][c])
		dfs(r-1, c, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
	}

	// Perform DFS from Pacific and Atlantic borders
	for c := 0; c < cols; c++ {
		dfs(0, c, pac, heights[0][c])          // Top row (Pacific)
		dfs(rows-1, c, atl, heights[rows-1][c]) // Bottom row (Atlantic)
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, pac, heights[r][0])          // Left column (Pacific)
		dfs(r, cols-1, atl, heights[r][cols-1]) // Right column (Atlantic)
	}

	// Collect common coordinates
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			key := [2]int{r, c}
			if pac[key] && atl[key] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
