package dp2d

// Leetcode #518
func change(amount int, coins []int) int {
	r := make([]int, amount+1)
	r[0] = 1

	for i := len(coins) - 1; i >= 0; i-- {

		nr := make([]int, amount+1)
		nr[0] = 1

		for a := 1; a < amount+1; a++ {
			nr[a] = r[a]
			if a-coins[i] >= 0 {
				nr[a] += nr[a-coins[i]]
			}
		}

		r = nr
	}

	return r[amount]
}