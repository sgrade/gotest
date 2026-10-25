// 1254. Number of Closed Islands
// https://leetcode.com/problems/number-of-closed-islands/

package numberofclosedislands

// closedIsland counts land components (0s) that never touch the grid border.
func closedIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

	// bfs explores the land component containing (startR, startC)
	// and reports whether it is closed (never reaches the border).
	bfs := func(startR, startC int) bool {
		queue := [][2]int{{startR, startC}}
		visited[startR][startC] = true
		isClosed := true

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				r, c := cur[0]+dir[0], cur[1]+dir[1]
				if r < 0 || r >= rows || c < 0 || c >= cols {
					// Stepping off the grid means the component touches the border.
					isClosed = false
				} else if grid[r][c] == 0 && !visited[r][c] {
					visited[r][c] = true
					queue = append(queue, [2]int{r, c})
				}
			}
		}

		return isClosed
	}

	closedIslands := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == 0 && !visited[r][c] && bfs(r, c) {
				closedIslands++
			}
		}
	}

	return closedIslands
}
