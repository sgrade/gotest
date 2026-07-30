// 3016. Minimum Number of Pushes to Type Word II
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/

package minimumnumberofpushestotypewordii

import "slices"

func minimumPushes(word string) int {
	var cntr [26]int
	for _, c := range word {
		cntr[c-'a']++
	}
	cnts := cntr[:]
	slices.Sort(cnts)
	slices.Reverse(cnts)
	ans := 0
	for i, cnt := range cnts {
		ans += (i/8 + 1) * cnt
	}
	return ans
}
