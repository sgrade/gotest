package templates

// Graph - Adjacency List
func buildGraph(edges [][]int, n int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // for undirected
	}
	return graph
}

// DFS - Recursive
func dfsRecursive(node int, graph [][]int, visited []bool) {
	visited[node] = true
	// process node
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfsRecursive(neighbor, graph, visited)
		}
	}
}

// DFS - Iterative (Stack)
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

// BFS - Queue
func bfs(start int, graph [][]int) {
	visited := make([]bool, len(graph))
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// process node
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// BFS - Level Order (Distance)
func bfsLevels(start int, graph [][]int) {
	visited := make([]bool, len(graph))
	queue := []int{start}
	visited[start] = true
	level := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// process node at level
			_ = level
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
		level++
	}
}

