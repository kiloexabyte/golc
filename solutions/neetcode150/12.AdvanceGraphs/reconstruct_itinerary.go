package advancegraphs

import "sort"

// Leetcode #332
func findItinerary(tickets [][]string) []string {
	// Sort tickets to guarantee lexical order
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})

	graph := make(map[string][]string)
	for _, ticket := range tickets {
		graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
	}

	// Use DFS post-order traversal (Hierholzer’s algorithm)
	var res []string
	var dfs func(string)
	dfs = func(airport string) {
		for len(graph[airport]) > 0 {
			// Always take the lex smallest destination
			next := graph[airport][0]
			graph[airport] = graph[airport][1:]
			dfs(next)
		}
		res = append(res, airport)
	}

	dfs("JFK")

	// Reverse result since we added airports post-order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
