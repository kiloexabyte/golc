package graphs

// Leetcode #200
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visit := map[[2]int]struct{}{}
	res := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		q := [][2]int{{r, c}}
		visit[[2]int{r, c}] = struct{}{}

		directions := [][2]int{
			{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		}

		for len(q) > 0 {
			current := q[0]
			q = q[1:] // pop front
			cr, cc := current[0], current[1]

			for _, d := range directions {
				newR, newC := cr+d[0], cc+d[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols {
					if grid[newR][newC] == '1' {
						if _, exists := visit[[2]int{newR, newC}]; !exists {
							q = append(q, [2]int{newR, newC})
							visit[[2]int{newR, newC}] = struct{}{}
						}
					}
				}
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				if _, exists := visit[[2]int{r, c}]; !exists {
					bfs(r, c)
					res++
				}
			}
		}
	}

	return res
}
