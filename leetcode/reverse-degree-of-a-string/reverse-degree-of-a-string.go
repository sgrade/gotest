// 3498. Reverse Degree of a String
// https://leetcode.com/problems/reverse-degree-of-a-string/

package reversedegreeofastring

func reverseDegree(s string) int {
	ans := 0
	for i, ch := range s {
		reversePos := 26 - int(ch-'a')
		ans += reversePos * (i + 1)
	}
	return ans
}
