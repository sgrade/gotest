// 1807. Evaluate the Bracket Pairs of a String
// https://leetcode.com/problems/evaluate-the-bracket-pairs-of-a-string/

package evaluatethebracketpairsofastring

import "strings"

func evaluate(s string, knowledge [][]string) string {
	lookup := make(map[string]string, len(knowledge))
	for _, kv := range knowledge {
		lookup[kv[0]] = kv[1]
	}

	var sb strings.Builder
	sb.Grow(len(s))
	// left is the start of the current key, or -1 when outside brackets.
	left := -1
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == '(':
			left = i + 1
		case c == ')':
			if val, ok := lookup[s[left:i]]; ok {
				sb.WriteString(val)
			} else {
				sb.WriteByte('?')
			}
			left = -1
		case left < 0:
			// Plain character outside any bracket pair.
			sb.WriteByte(c)
		}
	}
	return sb.String()
}
