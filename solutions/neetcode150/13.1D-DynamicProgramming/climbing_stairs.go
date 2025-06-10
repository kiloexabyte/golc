package dp

// Leetcode #70
func climbStairs(n int) int {
    var lastDigit = 0
	var k = 1
	var fib = 0

	for i := 0; i < n; i++ {
		var sum = k + lastDigit
		lastDigit = k
		k = sum
		fib = sum
	}

	return fib
}