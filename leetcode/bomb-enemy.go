// 361. Bomb Enemy
// https://leetcode.com/problems/bomb-enemy/

package leetcode

// Optimized with Leetcode's Python Sample 615 ms submission
func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	ans := 0

	dp := make([][]int, rows)
	for row := range dp {
		dp[row] = make([]int, cols)
	}

	for row_idx := range grid {
		enemies := 0
		empty_cells := []int{}
		for col_idx := 0; col_idx <= cols; col_idx++ {
			if col_idx == cols || grid[row_idx][col_idx] == 'W' {
				for _, idx := range empty_cells {
					dp[row_idx][idx] = enemies
				}
				enemies = 0
				empty_cells = empty_cells[:0]
			} else if grid[row_idx][col_idx] == 'E' {
				enemies++
			} else {
				empty_cells = append(empty_cells, col_idx)
			}
		}
	}

	for col_idx := 0; col_idx < cols; col_idx++ {
		enemies := 0
		empty_cells := []int{}
		for row_idx := 0; row_idx <= rows; row_idx++ {
			if row_idx == rows || grid[row_idx][col_idx] == 'W' {
				for _, idx := range empty_cells {
					dp[idx][col_idx] += enemies
					if dp[idx][col_idx] > ans {
						ans = dp[idx][col_idx]
					}
				}
				enemies = 0
				empty_cells = empty_cells[:0]
			} else if grid[row_idx][col_idx] == 'E' {
				enemies++
			} else {
				empty_cells = append(empty_cells, row_idx)
			}
		}
	}

	return ans
}
