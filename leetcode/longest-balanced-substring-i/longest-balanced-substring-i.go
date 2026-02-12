// 3713. Longest Balanced Substring I
//

package longestbalancedsubstringi

// Based on Editorial's Approach: Enumeration
func longestBalanced(s string) int {
	n := len(s)
	maxLen := 0
	for i := 0; i < n; i++ {
		counter := make([]int, 26)
		for j := i; j < n; j++ {
			idx := s[j] - 'a'
			counter[idx]++
			balanced := true
			for _, cnt := range counter {
				if cnt > 0 && cnt != counter[idx] {
					balanced = false
					break
				}
			}
			if balanced && (j-i+1) > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}
