package graphs
// Leetcode #695
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := map[[2]int]struct{}{}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		// Base cases
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
			return 0
		}
		if _, seen := visited[[2]int{r, c}]; seen {
			return 0
		}

		// Mark as visited
		visited[[2]int{r, c}] = struct{}{}

		// Explore all 4 directions
		return 1 +
			dfs(r+1, c) +
			dfs(r-1, c) +
			dfs(r, c+1) +
			dfs(r, c-1)
	}

	maxArea := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
