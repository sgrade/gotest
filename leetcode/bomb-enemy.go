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

	for rowIdx := range grid {
		enemies := 0
		emptyCells := []int{}
		for colIdx := 0; colIdx <= cols; colIdx++ {
			if colIdx == cols || grid[rowIdx][colIdx] == 'W' {
				for _, idx := range emptyCells {
					dp[rowIdx][idx] = enemies
				}
				enemies = 0
				emptyCells = emptyCells[:0]
			} else if grid[rowIdx][colIdx] == 'E' {
				enemies++
			} else {
				emptyCells = append(emptyCells, colIdx)
			}
		}
	}

	for colIdx := 0; colIdx < cols; colIdx++ {
		enemies := 0
		emptyCells := []int{}
		for rowIdx := 0; rowIdx <= rows; rowIdx++ {
			if rowIdx == rows || grid[rowIdx][colIdx] == 'W' {
				for _, idx := range emptyCells {
					dp[idx][colIdx] += enemies
					if dp[idx][colIdx] > ans {
						ans = dp[idx][colIdx]
					}
				}
				enemies = 0
				emptyCells = emptyCells[:0]
			} else if grid[rowIdx][colIdx] == 'E' {
				enemies++
			} else {
				emptyCells = append(emptyCells, rowIdx)
			}
		}
	}

	return ans
}
