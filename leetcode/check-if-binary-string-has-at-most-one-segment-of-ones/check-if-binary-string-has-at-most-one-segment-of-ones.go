// 1784. Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/

package checkifbinarystringhasatmostonesegmentofones

import "strings"

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
