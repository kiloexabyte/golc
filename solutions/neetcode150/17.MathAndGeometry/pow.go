package mathandgeometry

// Leetcode #50
func myPow(x float64, n int) float64 {
    N := int64(n)
    res := helper(x, abs(N))
    if N >= 0 {
        return res
    }
    return 1 / res
}

func helper(x float64, n int64) float64 {
    if x == 0 {
        return 0
    }
    if n == 0 {
        return 1
    }

    half := helper(x, n/2)
    result := half * half

    if n%2 != 0 {
        return x * result
    }
    return result
}

func abs(n int64) int64 {
    if n < 0 {
        return -n
    }
    return n
}
