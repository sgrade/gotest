// 1758. Minimum Changes To Make Alternating Binary String
// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/

package minimumchangestomakealternatingbinarystring

func minOperations(s string) int {
	first, second := byte('0'), byte('1')
	ansZero := 0
	for i := range s {
		if s[i] != first {
			ansZero++
		}
		first, second = second, first
	}
	return min(ansZero, len(s)-ansZero)
}
