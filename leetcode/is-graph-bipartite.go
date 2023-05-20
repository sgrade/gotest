// 785. Is Graph Bipartite?
// https://leetcode.com/problems/is-graph-bipartite/

package leetcode

func isBipartite(graph [][]int) bool {
	n := len(graph)
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = -1
	}

	for i := 0; i < n; i++ {
		if visited[i] != -1 {
			continue
		}
		visited[i] = 0
		var st Stack
		st.Push(i)
		for !st.Empty() {
			j := st.Pop()
			group_j := visited[j]
			for _, neigh := range graph[j] {
				if visited[neigh] == -1 {
					st.Push(neigh)
					visited[neigh] = group_j ^ 1
				} else if visited[neigh] == group_j {
					return false
				}
			}
		}
	}

	return true
}

type Stack []int

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
