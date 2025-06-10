package main

import "math"

// Leetcode #4
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2

	total := len(A) + len(B)
	half := total / 2

	// Ensure A is the smaller array
	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)

	for l <= r {
		i := (l + r) / 2
		j := half - i

		var ALeft, ARight, BLeft, BRight float64

		if i > 0 {
			ALeft = float64(A[i-1])
		} else {
			ALeft = math.Inf(-1)
		}

		if i < len(A) {
			ARight = float64(A[i])
		} else {
			ARight = math.Inf(1)
		}

		if j > 0 {
			BLeft = float64(B[j-1])
		} else {
			BLeft = math.Inf(-1)
		}

		if j < len(B) {
			BRight = float64(B[j])
		} else {
			BRight = math.Inf(1)
		}

		// Check partition correctness
		if ALeft <= BRight && BLeft <= ARight {
			if total%2 != 0 {
				return math.Min(ARight, BRight)
			}
			return (math.Max(ALeft, BLeft) + math.Min(ARight, BRight)) / 2
		} else if ALeft > BRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}

	panic("Input arrays are not valid.")
}