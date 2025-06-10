package dp

// leetcode #416
func canPartition(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum%2 == 1 {
        return false
    }
    target := sum / 2

    dp := map[int]struct{}{0: {}}

    for i := len(nums) - 1; i >= 0; i-- {
        nextDp := map[int]struct{}{}
        for t := range dp {
            if t+nums[i] == target {
                return true
            }
            nextDp[t+nums[i]] = struct{}{}
        }
        for k := range nextDp {
            dp[k] = struct{}{}
        }
    }

    _, exists := dp[target]
    return exists
}
