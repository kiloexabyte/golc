package graphs

// Leetcode #994
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	freshCount := 0
	queue := make([][2]int, 0)

	// Initialize queue with rotten oranges and count fresh ones
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				freshCount++
			} else if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}

	time := 0

	for len(queue) > 0 && freshCount > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:] // dequeue

			for _, d := range directions {
				newR, newC := r+d[0], c+d[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] == 1 {
					grid[newR][newC] = 2
					freshCount--
					queue = append(queue, [2]int{newR, newC})
				}
			}
		}
		time++
	}

	if freshCount == 0 {
		return time
	}
	return -1
}
