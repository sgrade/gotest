// 1653. Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/

package minimumdeletionstomakestringbalanced

func minimumDeletions(s string) int {
	minDeletions, bCount := 0, 0
	for i := range len(s) {
		if s[i] == 'b' {
			bCount++
		} else {
			minDeletions = min(minDeletions+1, bCount)
		}
	}
	return minDeletions
}
