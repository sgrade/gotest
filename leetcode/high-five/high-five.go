// 1086. High Five
// https://leetcode.com/problems/high-five/

package highfive

import "slices"

func highFive(items [][]int) [][]int {
	slices.SortFunc(items, func(a, b []int) int {
		return a[0] - b[0]
	})

	ans := make([][]int, 0)

	id := items[0][0]
	scores := make([]int, 0)
	for _, item := range items {
		curID, curScore := item[0], item[1]
		if curID == id {
			scores = append(scores, curScore)
		} else {
			slices.SortFunc(scores, func(a, b int) int {
				return b - a
			})
			curSum := 0
			for i := range min(len(scores), 5) {
				curSum += scores[i]
			}
			ans = append(ans, []int{id, curSum / 5})
			id = curID
			scores = []int{curScore}

		}
	}

	if len(scores) > 0 {
		slices.SortFunc(scores, func(a, b int) int {
			return b - a
		})
		curSum := 0
		for i := range min(len(scores), 5) {
			curSum += scores[i]
		}
		ans = append(ans, []int{id, curSum / 5})
	}

	return ans
}
