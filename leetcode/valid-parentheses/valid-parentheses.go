// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

package validparentheses

// Push opens (or when the stack is empty). Pop when the top matches the
// current close; otherwise push the mismatched close. Valid iff the stack
// ends empty.
func isValid(s string) bool {
	st := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' || len(st) == 0 {
			st = append(st, c)
		} else if st[len(st)-1] == '(' && c == ')' {
			st = st[:len(st)-1]
		} else if st[len(st)-1] == '[' && c == ']' {
			st = st[:len(st)-1]
		} else if st[len(st)-1] == '{' && c == '}' {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	if len(st) > 0 {
		return false
	}
	return true
}
