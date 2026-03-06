// 1784. Check if Binary String Has at Most One Segment of Ones
// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/

package checkifbinarystringhasatmostonesegmentofones

func checkOnesSegment(s string) bool {
	i := 0
	foundOne := false
	for i < len(s) {
		if s[i] == byte('1') {
			foundOne = true
			break
		}
		i++
	}

	for i < len(s) && s[i] == byte('1') {
		i++
	}
	for i < len(s) && s[i] == byte('0') {
		i++
	}

	for i < len(s) {
		if s[i] == '1' {
			return false
		}
	}

	if foundOne {
		return true
	}
	return false
}
