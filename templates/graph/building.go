package graph

// BuildGraph - Adjacency List
func buildGraph(edges [][]int, n int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // for undirected
	}
	return graph
}

