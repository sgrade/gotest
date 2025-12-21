// 955. Delete Columns to Make Sorted II
// https://leetcode.com/problems/delete-columns-to-make-sorted-ii/

package deletecolumnstomakesortedii

// Based on Editorial's Approach 2: Greedy with Optimizations
func minDeletionSize(strs []string) int {
	rows, cols := len(strs), len(strs[0])

	// cuts[r] tracks whether the relative ordering of rows r and r+1 has been established
	// (when they have a character difference in some column).
	established := make([]bool, rows-1)
	for i := range rows - 1 {
		established[i] = false
	}

	ans := 0
	for c := 0; c < cols; c++ {
		cont := false
		for r := 0; r < rows-1; r++ {
			if !established[r] && strs[r][c] > strs[r+1][c] {
				ans++
				cont = true
				break
			}
		}
		if cont {
			continue
		}

		for r := 0; r < rows-1; r++ {
			if strs[r][c] < strs[r+1][c] {
				established[r] = true
			}
		}
	}

	return ans
}
