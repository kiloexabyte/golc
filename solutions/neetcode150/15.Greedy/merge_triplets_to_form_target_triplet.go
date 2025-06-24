package greedy

// Leetcode #1899
func mergeTriplets(triplets [][]int, target []int) bool {
    good := map[int]struct{}{}

	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue
		}

		for i, v := range t {
			if v == target[i] {
				good[i] = struct{}{}
			}
		}
	}

	return len(good) == 3

}