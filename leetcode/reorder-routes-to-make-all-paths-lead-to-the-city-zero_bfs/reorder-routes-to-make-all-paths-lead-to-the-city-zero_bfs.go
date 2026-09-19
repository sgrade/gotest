// 1466. Reorder Routes to Make All Paths Lead to the City Zero
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

package reorderroutestomakeallpathsleadtothecityzerobfs

// Based on Editorial's Approach 2: Breadth First Search
func minReorder(n int, connections [][]int) int {
	type edge struct {
		to           int
		needsReorder bool // true if the original road points this way
	}

	adj := make([][]edge, n)
	for _, connection := range connections {
		from, to := connection[0], connection[1]
		adj[from] = append(adj[from], edge{to: to, needsReorder: true})
		adj[to] = append(adj[to], edge{to: from})
	}

	// Count roads pointing away from 0 on the unique path to 0.
	reorders := 0
	visited := make([]bool, n)
	visited[0] = true
	q := []int{0}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, e := range adj[node] {
			if visited[e.to] {
				continue
			}
			visited[e.to] = true
			if e.needsReorder {
				reorders++
			}
			q = append(q, e.to)
		}
	}
	return reorders
}
