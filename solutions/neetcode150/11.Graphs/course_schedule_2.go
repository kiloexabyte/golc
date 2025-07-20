package graphs

// Leetcode #210
func findOrder(numCourses int, prerequisites [][]int) []int {
	prereq := map[int][]int{}

	for i := 0; i < len(prerequisites); i++ {
		key := prerequisites[i][0]
		value := prerequisites[i][1]

		prereq[key] = append(prereq[key], value)
	}

	output := []int{}
	visit := map[int]struct{}{}
	cycle := map[int]struct{}{}

	var dfs func(crs int) bool
	dfs = func(crs int) bool {
		if _, exist := cycle[crs]; exist {
			return false
		}
		if _, exist := visit[crs]; exist {
			return true
		}

		cycle[crs] = struct{}{}
		for _, pre := range prereq[crs] {
			if !dfs(pre) {
				return false
			}
		}

		delete(cycle, crs)
		visit[crs] = struct{}{}
		output = append(output, crs)
		return true
	}

	for c := range numCourses {
		if !dfs(c) {
			return []int{}
		}
	}

	return output
}
