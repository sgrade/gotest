// 1351. Count Negative Numbers in a Sorted Matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

package countnegativenumbersinasortedmatrix

func countNegatives(grid [][]int) int {
	negatives := 0
	cols := len(grid[0])
	maxCol := cols
	for _, row := range grid {
		for col := 0; col < maxCol; col++ {
			if row[col] < 0 {
				negatives += cols - col
				maxCol = col + 1
				break
			}
		}
	}
	return negatives
}
