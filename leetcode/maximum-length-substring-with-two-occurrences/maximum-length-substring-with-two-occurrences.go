// 3090. Maximum Length Substring With Two Occurrences
// https://leetcode.com/problems/maximum-length-substring-with-two-occurrences/

package maximumlengthsubstringwithtwooccurrences

// Sliding window: keep [left, right] valid by shrinking from the left
// whenever the letter just added occurs more than twice.
func maximumLengthSubstring(s string) int {
	var count ['z' - 'a' + 1]int
	ans, left := 0, 0
	for right := 0; right < len(s); right++ {
		count[s[right]-'a']++
		for count[s[right]-'a'] > 2 {
			count[s[left]-'a']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
