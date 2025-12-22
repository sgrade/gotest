// 960. Delete Columns to Make Sorted III
// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/

package deletecolumnstomakesortediii

// Based on Editorial's Approach 1: Dynamic Programming
func minDeletionSize(strs []string) int {
	cols := len(strs[0])
	dp := make([]int, cols)
	for i := range dp {
		dp[i] = 1
	}

	for leftCol := cols - 2; leftCol >= 0; leftCol-- {
		for rightCol := leftCol + 1; rightCol < cols; rightCol++ {
			cont := false
			for _, row := range strs {
				if row[leftCol] > row[rightCol] {
					cont = true
					break
				}
			}
			if cont {
				continue
			}
			dp[leftCol] = max(dp[leftCol], 1+dp[rightCol])
		}
	}

	colsToKeep := 0
	for _, curAns := range dp {
		colsToKeep = max(colsToKeep, curAns)
	}
	return cols - colsToKeep
}
