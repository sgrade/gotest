// 1466. Reorder Routes to Make All Paths Lead to the City Zero
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

package reorderroutestomakeallpathsleadtothecityzero

// Based on Editorial's Approach 1: Depth First Search
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

	var dfs func(int, int) int
	dfs = func(node, parent int) int {
		// Count roads pointing away from 0 on the unique path to 0.
		reorders := 0
		for _, e := range adj[node] {
			if e.to == parent {
				continue
			}
			if e.needsReorder {
				reorders++
			}
			reorders += dfs(e.to, node)
		}
		return reorders
	}

	return dfs(0, -1)
}
