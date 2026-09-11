// 1086. High Five
// https://leetcode.com/problems/high-five/

package highfive

import (
	"cmp"
	"slices"
)

func highFive(items [][]int) [][]int {
	// Group each student's scores together by sorting on id.
	slices.SortFunc(items, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	var ans [][]int
	id := items[0][0]
	var scores []int
	for _, item := range items {
		curID, curScore := item[0], item[1]
		if curID == id {
			scores = append(scores, curScore)
			continue
		}
		ans = append(ans, []int{id, topFiveAverage(scores)})
		id = curID
		scores = []int{curScore}
	}
	if len(scores) > 0 {
		ans = append(ans, []int{id, topFiveAverage(scores)})
	}
	return ans
}

// topFiveAverage returns the integer average of the five highest scores.
func topFiveAverage(scores []int) int {
	slices.SortFunc(scores, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	sum := 0
	for i := range min(len(scores), 5) {
		sum += scores[i]
	}
	return sum / 5
}
