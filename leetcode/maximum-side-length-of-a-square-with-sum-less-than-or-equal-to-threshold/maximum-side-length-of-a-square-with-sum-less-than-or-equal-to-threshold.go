// 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/

package maximumsidelengthofasquarewithsumlessthanorequaltothreshold

// Based on Editorial's Approach 1: Binary Search
func maxSideLength(mat [][]int, threshold int) int {
	rows, cols := len(mat), len(mat[0])

	// Build prefix sum array for efficient range sum queries
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + mat[i-1][j-1]
		}
	}

	// Binary search on the side length
	left, right, maxLen := 1, min(rows, cols), 0
	for left <= right {
		mid := (left + right) / 2
		found := false

		// Check if there exists a square of side length mid with sum <= threshold
		for i := 1; i <= rows-mid+1; i++ {
			for j := 1; j <= cols-mid+1; j++ {
				squareSum := prefixSum[i+mid-1][j+mid-1] - prefixSum[i-1][j+mid-1] - prefixSum[i+mid-1][j-1] + prefixSum[i-1][j-1]
				if squareSum <= threshold {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if found {
			maxLen = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return maxLen
}
