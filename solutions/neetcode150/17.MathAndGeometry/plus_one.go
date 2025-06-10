package mathandgeometry

// Leetcode #66
func plusOne(digits []int) []int {
    n := len(digits)
    digits[n-1]++
    i := 1

    for digits[n-i] > 9 {
        digits[n-i] = 0
        if i < n {
            digits[n-i-1]++
        } else {
            // Insert 1 at the beginning
            digits = append([]int{1}, digits...)
            break
        }
        i++
    }

    return digits
}
