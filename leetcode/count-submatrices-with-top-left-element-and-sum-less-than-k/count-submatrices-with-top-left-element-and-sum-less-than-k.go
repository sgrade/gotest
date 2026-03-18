// 3070. Count Submatrices with Top-Left Element and Sum Less Than k
// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/

package countsubmatriceswithtopleftelementandsumlessthank

// Based on Editorial's Approach: 2D Prefix Sum
//
// colSum[j] accumulates the column prefix sum up to the current row.
// Adding colSum values left-to-right gives the 2D prefix sum for the
// submatrix (0,0)→(i,j), counted when ≤ k.
func countSubmatrices(grid [][]int, k int) int {
	colSum := make([]int, len(grid[0]))
	count := 0

	for _, row := range grid {
		prefixSum := 0
		for j, val := range row {
			colSum[j] += val
			prefixSum += colSum[j]
			if prefixSum <= k {
				count++
			}
		}
	}

	return count
}
