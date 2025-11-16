package numberofsubstringswithonly1s

// 1513. Number of Substrings With Only 1s
// https://leetcode.com/problems/number-of-substrings-with-only-1s/

func numSub(s string) int {
	MOD := int(1e9 + 7)
	ans, n, ones := 0, len(s), 0
	for i := range n {
		if s[i] == '0' {
			ones = 0
			continue
		}
		ones++
		ans += ones
		ans %= MOD
	}
	return ans
}
