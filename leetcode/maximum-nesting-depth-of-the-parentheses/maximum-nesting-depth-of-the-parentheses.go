// 1614. Maximum Nesting Depth of the Parentheses
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/

package maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	// depth tracks the number of currently open parentheses;
	// ans records the deepest level reached.
	ans, depth := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			depth++
			ans = max(ans, depth)
		case ')':
			depth--
		}
	}
	return ans
}
