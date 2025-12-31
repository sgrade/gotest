package graph

// DFSRecursive - Recursive DFS
func dfsRecursive(node int, graph [][]int, visited []bool) {
	visited[node] = true
	// process node
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfsRecursive(neighbor, graph, visited)
		}
	}
}

// DFSIterative - Iterative DFS using Stack
func dfsIterative(start int, graph [][]int) {
	visited := make([]bool, len(graph))
	stack := []int{start}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[node] {
			continue
		}
		visited[node] = true
		// process node
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
}

