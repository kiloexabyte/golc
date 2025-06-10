package main

import "strings"

// Leetcode #51
func solveNQueens(n int) [][]string {
	cols := make(map[int]bool)
	posDiag := make(map[int]bool)
	negDiag := make(map[int]bool)
	res := make([][]string,0)
	board := make([][]string, 0)

	var backtrack func (r, n int)
	backtrack = func (r, n int) {
		if r == n {
			copy := []string{}
			for _, row := range board {
				copy = append(copy, strings.Join(row, ""))
			}
			res = append(res, copy)
			return
		}

		for c := 0; c < n; c++ {
			if cols[c] || posDiag[r+c] || negDiag[r-c] {
				continue
			}
			cols[c] = true
			posDiag[r + c] = true
			negDiag[r - c] = true
			board[r][c] = "Q"

			backtrack(r + 1, n)

			cols[c] = false
			posDiag[r + c] = false
			negDiag[r - c] = false
			board[r][c] = "."
		}
	}

	for i := 0; i < n; i++ {
		row := []string{}
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}

	backtrack(0, n)
	return res
}
