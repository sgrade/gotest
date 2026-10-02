// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/

package generateparentheses

func generateParenthesis(n int) []string {
	var ans []string

	// backtrack builds every valid string by appending one parenthesis at a time.
	var backtrack func(s string, openCount, closeCount int)
	backtrack = func(s string, openCount, closeCount int) {
		if len(s) == n*2 {
			ans = append(ans, s)
			return
		}
		// An opening parenthesis is allowed while fewer than n are used.
		if openCount < n {
			backtrack(s+"(", openCount+1, closeCount)
		}
		// A closing parenthesis is allowed only if it matches an unclosed one.
		if closeCount < openCount {
			backtrack(s+")", openCount, closeCount+1)
		}
	}

	backtrack("", 0, 0)
	return ans
}
