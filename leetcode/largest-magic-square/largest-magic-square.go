package largestmagicsquare

// 1895. Largest Magic Square
// https://leetcode.com/problems/largest-magic-square/

// Based on Editorial's Approach: Enumerating Squares + Prefix Sum Optimization
func largestMagicSquare(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// Build row prefix sums
	prefixRow := make([][]int, rows)
	for i := range prefixRow {
		prefixRow[i] = make([]int, cols)
		prefixRow[i][0] = grid[i][0]
		for j := 1; j < cols; j++ {
			prefixRow[i][j] = prefixRow[i][j-1] + grid[i][j]
		}
	}

	// Build column prefix sums
	prefixCol := make([][]int, rows)
	for i := range prefixCol {
		prefixCol[i] = make([]int, cols)
	}
	for j := 0; j < cols; j++ {
		prefixCol[0][j] = grid[0][j]
		for i := 1; i < rows; i++ {
			prefixCol[i][j] = prefixCol[i-1][j] + grid[i][j]
		}
	}

	for side := min(rows, cols); side >= 2; side-- {
		for i := 0; i+side <= rows; i++ {
			for j := 0; j+side <= cols; j++ {
				// Calculate target sum from first row
				targetSum := prefixRow[i][j+side-1]
				if j > 0 {
					targetSum -= prefixRow[i][j-1]
				}

				// Check all rows have same sum
				valid := true
				for r := i + 1; r < i+side; r++ {
					rowSum := prefixRow[r][j+side-1]
					if j > 0 {
						rowSum -= prefixRow[r][j-1]
					}
					if rowSum != targetSum {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}

				// Check all columns have same sum
				for c := j; c < j+side; c++ {
					colSum := prefixCol[i+side-1][c]
					if i > 0 {
						colSum -= prefixCol[i-1][c]
					}
					if colSum != targetSum {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}

				// Check diagonals
				diagonal, antiDiagonal := 0, 0
				for k := 0; k < side; k++ {
					diagonal += grid[i+k][j+k]
					antiDiagonal += grid[i+k][j+side-1-k]
				}
				if diagonal == targetSum && antiDiagonal == targetSum {
					return side
				}
			}
		}
	}
	return 1
}
