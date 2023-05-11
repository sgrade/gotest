// 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/

package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	sIdx, gIdx, ans := len(s)-1, len(g)-1, 0
	for sIdx >= 0 && gIdx >= 0 {
		if s[sIdx] >= g[gIdx] {
			sIdx--
			ans++
		}
		gIdx--
	}
	return ans
}
