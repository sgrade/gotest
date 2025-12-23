// 2054. Two Best Non-Overlapping Events
// https://leetcode.com/problems/two-best-non-overlapping-events/

package twobestnonoverlappingevents

import "slices"

// Based on Editorial's Approach 3: Greedy
func maxTwoEvents(events [][]int) int {
	times := make([][]int, 0)
	for _, e := range events {
		startTime := []int{e[0], 1, e[2]}
		times = append(times, startTime)
		endTime := []int{e[1] + 1, 0, e[2]}
		times = append(times, endTime)
	}
	slices.SortFunc(times, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	ans, maxValSoFar := 0, 0
	for _, t := range times {
		if t[1] == 1 {
			ans = max(ans, t[2]+maxValSoFar)
		} else {
			maxValSoFar = max(maxValSoFar, t[2])
		}
	}
	return ans
}
