// 1415. The k-th Lexicographical String of All Happy Strings of Length n
// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/

package thekthlexicographicalstringofallhappystringsoflengthn

// Based on Editorial's Approach 3: Iterative Using a Stack
// Iterative DFS using an explicit stack. Pushes children in reverse order
// ('c'..'a') so that popping yields lexicographic order. Counts complete
// strings of length n until the k-th is found.
func getHappyString(n int, k int) string {
	stack := []string{""}
	idx := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(cur) == n {
			idx++
			if idx == k {
				return cur
			}
			continue
		}

		// Push 'c'..'a' so 'a' is on top and explored first.
		for ch := byte('c'); ch >= 'a'; ch-- {
			if len(cur) > 0 && cur[len(cur)-1] == ch {
				continue
			}
			stack = append(stack, cur+string(ch))
		}
	}

	return ""
}
