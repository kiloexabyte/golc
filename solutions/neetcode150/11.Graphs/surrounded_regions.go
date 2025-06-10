package graphs

// Leetcode #130
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// DFS helper to mark boundary-connected 'O's as 'T'
	var capture func(r, c int)
	capture = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T'
		capture(r+1, c)
		capture(r-1, c)
		capture(r, c+1)
		capture(r, c-1)
	}

	// Step 1: Mark 'O's on the boundary and their connected 'O's as 'T'
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			capture(i, 0)
		}
		if board[i][cols-1] == 'O' {
			capture(i, cols-1)
		}
	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			capture(0, j)
		}
		if board[rows-1][j] == 'O' {
			capture(rows-1, j)
		}
	}

	// Step 2: Convert all unmarked 'O's to 'X', and revert 'T's back to 'O'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}


