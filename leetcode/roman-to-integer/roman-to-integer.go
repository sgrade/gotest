// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/

package romantointeger

func romanToInt(s string) int {
	chars := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	current := 0
	prev := 0
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		current = chars[string(s[i])]
		if current < prev {
			ans -= current
		} else {
			ans += current
			prev = current
		}
	}
	return ans
}
