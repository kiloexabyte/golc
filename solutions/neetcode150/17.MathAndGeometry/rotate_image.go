package mathandgeometry

// Leetcode #48
func rotate(matrix [][]int) {
    l := 0
    r := len(matrix) - 1

    for l < r {
        for i := 0; i < r-l; i++ {
            top := l
            bottom := r

            // save top-left
            topLeft := matrix[top][l+i]

            // move bottom-left to top-left
            matrix[top][l+i] = matrix[bottom-i][l]

            // move bottom-right to bottom-left
            matrix[bottom-i][l] = matrix[bottom][r-i]

            // move top-right to bottom-right
            matrix[bottom][r-i] = matrix[top+i][r]

            // move saved top-left to top-right
            matrix[top+i][r] = topLeft
        }

        l++
        r--
    }
}
