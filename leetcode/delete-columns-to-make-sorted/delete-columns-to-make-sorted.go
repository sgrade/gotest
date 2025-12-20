// 944. Delete Columns to Make Sorted
// https://leetcode.com/problems/delete-columns-to-make-sorted/

package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	rows, cols := len(strs), len(strs[0])
	ans := 0
	for c := range cols {
		for r := 1; r < rows; r++ {
			if strs[r-1][c] > strs[r][c] {
				ans++
				break
			}
		}
	}
	return ans
}
