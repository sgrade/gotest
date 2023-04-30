// 1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

package leetcode

func longestSubarray(nums []int, limit int) int {
	minq, maxq := []int{}, []int{}
	left := 0
	right := 0
	for ; right < len(nums); right++ {
		for len(minq) > 0 && minq[len(minq)-1] > nums[right] {
			minq = minq[:len(minq)-1]
		}
		minq = append(minq, nums[right])
		for len(maxq) > 0 && maxq[len(maxq)-1] < nums[right] {
			maxq = maxq[:len(maxq)-1]
		}
		maxq = append(maxq, nums[right])
		if maxq[0]-minq[0] > limit {
			if minq[0] == nums[left] {
				minq = minq[1:]
			}
			if maxq[0] == nums[left] {
				maxq = maxq[1:]
			}
			left++
		}
	}
	return right - left
}
