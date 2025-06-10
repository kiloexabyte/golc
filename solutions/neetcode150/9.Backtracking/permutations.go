package main

// Leetcode #46
func permute(nums []int) [][]int {
    res := [][]int{}

    if len(nums) == 1 {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        res = append(res, tmp)
        return res
    }

    for i := 0; i < len(nums); i++ {
        n := nums[0]
        nums = nums[1:] // remove the first element

        ps := permute(nums)

        for _, p := range ps {
            newPerm := append([]int{}, p...) // copy p to avoid mutation
            newPerm = append(newPerm, n)
            res = append(res, newPerm)
        }

        nums = append(nums, n) // add the element back for the next round
    }

    return res
}
