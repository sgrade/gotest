package minimumdifferencebetweenhighestandlowestofkscores

import "slices"

func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)
	minDiff := int(1e5 + 1)
	for i := k; i <= len(nums); i++ {
		curDiff := nums[i-1] - nums[i-k]
		minDiff = min(minDiff, curDiff)
	}
	return minDiff
}
