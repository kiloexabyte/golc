package bitmanipulation

// Leetcode #7
func reverse(x int) int {
    min := -2147483648
	max := 2147483648

	res := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if res > max / 10 || (res == max / 10 && digit >= max % 10) {
			return 0
		}
		if res < min / 10 || (res == min / 10 && digit <= min % 10) {
			return 0
		}

		res = (res * 10) + digit
	}

	return res
}