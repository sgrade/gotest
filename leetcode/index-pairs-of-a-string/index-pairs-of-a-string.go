// 1065. Index Pairs of a String
// https://leetcode.com/problems/index-pairs-of-a-string/

package indexpairsofastring

import (
	"sort"
	"strings"
)

func indexPairs(text string, words []string) [][]int {
	ans := [][]int{}
	for _, w := range words {
		idx, prev := 0, 0
		for {
			idx = strings.Index(text[idx:], w)
			if idx == -1 {
				break
			}
			ans = append(ans, []int{prev + idx, prev + idx + len(w) - 1})
			prev = prev + idx + 1
			idx = prev
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] == ans[j][0] {
			return ans[i][1] < ans[j][1]
		} else {
			return ans[i][0] < ans[j][0]
		}
	})
	return ans
}
