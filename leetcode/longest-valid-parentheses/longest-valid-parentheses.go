// 32. Longest Valid Parentheses
// https://leetcode.com/problems/longest-valid-parentheses/

// Based on Editorial's Approach 3: Using Stack

package longestvalidparentheses

func longestValidParentheses(s string) int {
	longest := 0

	// The stack holds indices of unmatched '(' and, at the bottom, the index
	// just before the current valid substring (initially -1).
	stack := []int{-1}
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
			continue
		}

		// Match the ')' with the most recent unmatched '('.
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			// Unmatched ')': it becomes the new base for later substrings.
			stack = append(stack, i)
			continue
		}
		// The valid substring extends from just after the stack top to i.
		longest = max(longest, i-stack[len(stack)-1])
	}
	return longest
}
