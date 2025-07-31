package advancegraphs

import "math"

// Leetcode 787
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prices := make([]int, n)
    for i := range prices {
        prices[i] = math.MaxInt32
    }
    prices[src] = 0

	for i := 0; i < k+1; i++ {
		tempPrices := make([]int, len(prices))
		copy(tempPrices, prices)
		for _, flight := range flights {
			s := flight[0] // source
			d := flight[1] // destination
			p := flight[2] // price
			if prices[s] == math.MaxInt32 {
				continue
			}
			if prices[s] + p < tempPrices[d] {
				tempPrices[d] = prices[s] + p
			}
		}

		prices = tempPrices
	}

	if prices[dst] == math.MaxInt32 {
		return -1 
	}
	return prices[dst]
}