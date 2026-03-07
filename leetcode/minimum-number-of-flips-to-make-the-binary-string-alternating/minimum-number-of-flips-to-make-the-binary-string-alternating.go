// 1888. Minimum Number of Flips to Make the Binary String Alternating
// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/

package minimumnumberofflipstomakethebinarystringalternating

// Based on https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/solutions/7631262/solution-by-la_castille-yiau
func minFlips(s string) int {
	n := len(s)
	var op [2]int

	for i := range n {
		op[int(s[i]-'0')^(i&1)]++
	}

	res := min(op[0], op[1])

	if n%2 == 1 {
		for i := range n {
			b := int(s[i]-'0') ^ (i & 1)
			op[b]--
			op[1-b]++
			res = min(res, min(op[0], op[1]))
		}
	}

	return res
}
