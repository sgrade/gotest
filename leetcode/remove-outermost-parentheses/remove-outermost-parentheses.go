// 1021. Remove Outermost Parentheses
// https://leetcode.com/problems/remove-outermost-parentheses/

package removeoutermostparentheses

func removeOuterParentheses(s string) string {
	ans := ""
	stack := []rune{}
	for _, c := range s {
		if c == ')' {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans += string(c)
		}
		if c == '(' {
			stack = append(stack, c)
		}
	}
	return ans
}
