package templates

// Stack - Basic Operations
func stackBasics() {
	stack := []int{}
	// push
	stack = append(stack, 1)
	// pop
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		_ = top
	}
	// peek
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		_ = top
	}
}

// Monotonic Stack - Next Greater Element
func nextGreater(arr []int) []int {
	result := make([]int, len(arr))
	stack := []int{} // stores indices
	for i := len(arr) - 1; i >= 0; i-- {
		// pop smaller elements
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = arr[stack[len(stack)-1]]
		}
		stack = append(stack, i)
	}
	return result
}

// Valid Parentheses Pattern
func validParentheses(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

