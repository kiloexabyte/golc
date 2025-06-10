package dp

// Leetcode #746
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)

	for i := len(cost) - 3; i > -1; i-- {
		cost[i] += Min(cost[i+1], cost[i+2])
	}

	return Min(cost[0], cost[1])
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
