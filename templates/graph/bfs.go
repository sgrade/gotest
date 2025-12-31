package graph

// BFS - Queue-based BFS
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

// BFSLevels - BFS with Level Order (Distance)
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

