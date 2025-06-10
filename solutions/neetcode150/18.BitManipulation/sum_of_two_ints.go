package bitmanipulation

// Leetcode #371
func getSum(a int, b int) int {
    for b != 0 {
        tmp := (a & b) << 1
        a = a ^ b
        b = tmp
    }

    return a
}