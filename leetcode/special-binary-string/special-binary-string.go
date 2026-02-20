// 761. Special Binary String
// https://leetcode.com/problems/special-binary-string/

package specialbinarystring

import (
	"slices"
	"strings"
)

// Based on Editorial's Approach 1: Recursion
func makeLargestSpecial(s string) string {
	var specials []string
	balance, anchor := 0, 0
	for i := range s {
		if s[i] == '1' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			inner := makeLargestSpecial(s[anchor+1 : i])
			specials = append(specials, "1"+inner+"0")
			anchor = i + 1
		}
	}
	slices.SortFunc(specials, func(a, b string) int {
		return strings.Compare(b, a)
	})
	return strings.Join(specials, "")
}
