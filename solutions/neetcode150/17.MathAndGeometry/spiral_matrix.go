package mathandgeometry

// Leetcode #54
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }

    top, bottom := 0, len(matrix)
    left, right := 0, len(matrix[0])
    res := []int{}

    for left < right && top < bottom {
        // Traverse top row
        for i := left; i < right; i++ {
            res = append(res, matrix[top][i])
        }
        top++

        // Traverse right column
        for i := top; i < bottom; i++ {
            res = append(res, matrix[i][right-1])
        }
        right--

        // Check if boundaries still valid
        if !(left < right && top < bottom) {
            break
        }

        // Traverse bottom row
        for i := right - 1; i >= left; i-- {
            res = append(res, matrix[bottom-1][i])
        }
        bottom--

        // Traverse left column
        for i := bottom - 1; i >= top; i-- {
            res = append(res, matrix[i][left])
        }
        left++
    }

    return res
}
