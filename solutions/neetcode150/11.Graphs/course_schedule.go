package graphs

// Leetcode #207
func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		preMap[i] = []int{}
	}

	for _, pair := range prerequisites {
		course, prereq := pair[0], pair[1]
		preMap[course] = append(preMap[course], prereq)
	}

	visitSet := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(crs int) bool {
		if visitSet[crs] {
			return false // cycle detected
		}
		if len(preMap[crs]) == 0 {
			return true // no prerequisites
		}

		visitSet[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		visitSet[crs] = false
		preMap[crs] = []int{} // memoize as no-cycle

		return true
	}

	for crs := 0; crs < numCourses; crs++ {
		if !dfs(crs) {
			return false
		}
	}

	return true
}