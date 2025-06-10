package main

func isValidSudoku(board [][]byte) bool {
    cols := make(map[int][]byte)
    rows := make(map[int][]byte)
    squares := make(map[[2]int][]byte)

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }

            // Initialize row if not present
            if _, exists := rows[r]; !exists {
                rows[r] = []byte{}
            }

            // Initialize column if not present
            if _, exists := cols[c]; !exists {
                cols[c] = []byte{}
            }

            // Initialize square if not present
            squareKey := [2]int{r / 3, c / 3}
            if _, exists := squares[squareKey]; !exists {
                squares[squareKey] = []byte{}
            }

            // Check for duplicates in row, column, and square
            if contains(rows[r], board[r][c]) || 
               contains(cols[c], board[r][c]) || 
               contains(squares[squareKey], board[r][c]) {
                return false
            }

            // Add the number to the row, column, and square
            rows[r] = append(rows[r], board[r][c])
            cols[c] = append(cols[c], board[r][c])
            squares[squareKey] = append(squares[squareKey], board[r][c])
        }
    }

    return true
}

func contains(slice []byte, element byte) bool {
    for _, v := range slice {
        if v == element {
            return true
        }
    }
    return false
}