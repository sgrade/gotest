// 1200. Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/

package minimumabsolutedifference

import "slices"

func minimumAbsDifference(arr []int) [][]int {
	slices.Sort(arr)
	minDiff := int(1e5 + 1)
	for i := 1; i < len(arr); i++ {
		curDiff := arr[i] - arr[i-1]
		minDiff = min(minDiff, curDiff)
	}
	minDiffPairs := make([][]int, 0)
	for i := 1; i < len(arr); i++ {
		curDiff := arr[i] - arr[i-1]
		if curDiff == minDiff {
			curPair := []int{arr[i-1], arr[i]}
			minDiffPairs = append(minDiffPairs, curPair)
		}
	}
	return minDiffPairs
}
