// 463. Island Perimeter
// https://leetcode.com/problems/island-perimeter/

package leetcode

func islandPerimeter(grid [][]int) int {
	ans, rows, cols := 0, len(grid), len(grid[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				ans += 4
				if row > 0 && grid[row-1][col] == 1 {
					ans -= 2
				}
				if col > 0 && grid[row][col-1] == 1 {
					ans -= 2
				}
			}
		}
	}
	return ans
}
