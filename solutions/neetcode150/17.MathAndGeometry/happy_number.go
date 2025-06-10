package mathandgeometry

// Leetcode #202
func isHappy(n int) bool {
    seen := make(map[int]bool)

    sumOfSquares := func(num int) int {
        sum := 0
        for num > 0 {
            digit := num % 10
            sum += digit * digit
            num /= 10
        }
        return sum
    }

    newSum := sumOfSquares(n)

    for newSum != 1 {
        if seen[newSum] {
            return false
        }
        seen[newSum] = true
        newSum = sumOfSquares(newSum)
    }

    return true
}
