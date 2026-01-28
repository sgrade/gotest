// 3651. Minimum Cost Path with Teleportations
// https://leetcode.com/problems/minimum-cost-path-with-teleportations/

package minimumcostpathwithteleportations

import (
	"math"
	"sort"
)

// Based on Editorial's Approach: Dynamic Programming
func minCost(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])
	type point struct {
		row, col int
	}
	points := make([]point, 0, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			points = append(points, point{r, c})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return grid[points[i].row][points[i].col] < grid[points[j].row][points[j].col]
	})

	costs := make([][]int, rows)
	for r := range costs {
		costs[r] = make([]int, cols)
		for c := range costs[r] {
			costs[r][c] = math.MaxInt
		}
	}
	for teleportation := 0; teleportation <= k; teleportation++ {
		minCost := math.MaxInt
		for i, groupStart := 0, 0; i < len(points); i++ {
			minCost = min(minCost, costs[points[i].row][points[i].col])
			if i+1 < len(points) && grid[points[i].row][points[i].col] == grid[points[i+1].row][points[i+1].col] {
				continue
			}
			for j := groupStart; j <= i; j++ {
				costs[points[j].row][points[j].col] = minCost
			}
			groupStart = i + 1
		}
		for r := rows - 1; r >= 0; r-- {
			for c := cols - 1; c >= 0; c-- {
				if r == rows-1 && c == cols-1 {
					costs[r][c] = 0
					continue
				}
				if r != rows-1 {
					costs[r][c] = min(costs[r][c], costs[r+1][c]+grid[r+1][c])
				}
				if c != cols-1 {
					costs[r][c] = min(costs[r][c], costs[r][c+1]+grid[r][c+1])
				}
			}
		}
	}

	return costs[0][0]
}
