package mathandgeometry

// Leetcode #73
func setZeroes(matrix [][]int) {
    height := len(matrix)
    if height == 0 {
        return
    }
    width := len(matrix[0])

    zeroRows := make(map[int]bool)
    zeroCols := make(map[int]bool)

    // Identify rows and columns that need to be zeroed
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            if matrix[y][x] == 0 {
                zeroRows[y] = true
                zeroCols[x] = true
            }
        }
    }

    // Zero identified rows
    for row := range zeroRows {
        for i := 0; i < width; i++ {
            matrix[row][i] = 0
        }
    }

    // Zero identified columns
    for col := range zeroCols {
        for i := 0; i < height; i++ {
            matrix[i][col] = 0
        }
    }
}
