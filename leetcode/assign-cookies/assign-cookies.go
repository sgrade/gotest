// 455. Assign Cookies
// https://leetcode.com/problems/assign-cookies/

package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	sIdx, gIdx, ans := 0, 0, 0
	lenS, lenG := len(s), len(g)
	for sIdx < lenS && gIdx < lenG {
		if s[sIdx] >= g[gIdx] {
			gIdx++
			ans++
		}
		sIdx++
	}
	return ans
}
