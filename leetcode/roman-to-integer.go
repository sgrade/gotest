// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/

package leetcode

func romanToInt(s string) int {
	current, prev, ans := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch ch := s[i]; ch {
		case 'M':
			current = 1000
		case 'D':
			current = 500
		case 'C':
			current = 100
		case 'L':
			current = 50
		case 'X':
			current = 10
		case 'V':
			current = 5
		case 'I':
			current = 1
		}
		if current < prev {
			ans -= current
		} else {
			ans += current
		}
		prev = current
	}
	return ans
}
