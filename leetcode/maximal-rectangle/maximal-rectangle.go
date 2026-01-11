// 85. Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/

package maximalrectangle

// Based on Editorial's Approach 2: Dynamic Programming - Better Brute Force on Histograms
func maximalRectangle(matrix [][]byte) int {
	ans := 0
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == '1' {
				if col == 0 {
					dp[row][col] = 1
				} else {
					dp[row][col] = dp[row][col-1] + 1
				}
				width := dp[row][col]
				for r := row; r >= 0; r-- {
					width = min(width, dp[r][col])
					ans = max(ans, width*(row-r+1))
				}
			}
		}
	}

	return ans
}
