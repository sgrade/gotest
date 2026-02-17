// 696. Count Binary Substrings
// https://leetcode.com/problems/count-binary-substrings/

package countbinarysubstrings

func countBinarySubstrings(s string) int {
	substrings, prevCnt, curCnt := 0, 0, 1
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			substrings += min(prevCnt, curCnt)
			prevCnt = curCnt
			curCnt = 1
		} else {
			curCnt++
		}
	}
	return substrings + min(prevCnt, curCnt)
}
