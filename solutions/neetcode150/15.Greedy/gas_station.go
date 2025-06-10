package greedy

// Leetcode #134
func canCompleteCircuit(gas []int, cost []int) int {
    total := 0
    res := 0

    sumGas, sumCost := 0, 0
    for i := 0; i < len(gas); i++ {
        sumGas += gas[i]
        sumCost += cost[i]
    }
    if sumGas < sumCost {
        return -1
    }

    for i := 0; i < len(gas); i++ {
        total += gas[i] - cost[i]
        if total < 0 {
            total = 0
            res = i + 1
        }
    }

    return res
}
