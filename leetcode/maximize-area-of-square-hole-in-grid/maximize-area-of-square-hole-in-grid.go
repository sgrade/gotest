// 3454. Maximize Area of Square Hole in Grid
// https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/

package maximizeareaofsquareholeingrid

import "slices"

// longestConsecutive finds the longest sequence of consecutive integers in a sorted slice.
func longestConsecutive(bars []int) int {
	if len(bars) == 0 {
		return 0
	}
	maxLen, currentLen := 1, 1
	for i := 1; i < len(bars); i++ {
		if bars[i] == bars[i-1]+1 {
			currentLen++
		} else {
			currentLen = 1
		}
		maxLen = max(maxLen, currentLen)
	}
	return maxLen
}

// Find the longest consecutive sequence of bars in each direction.
// The square hole side is limited by the shorter sequence.
// Time: O(n log n + m log m), Space: O(1)
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	slices.Sort(hBars)
	slices.Sort(vBars)

	maxHConsecutive := longestConsecutive(hBars)
	maxVConsecutive := longestConsecutive(vBars)

	side := min(maxHConsecutive, maxVConsecutive) + 1
	return side * side
}
