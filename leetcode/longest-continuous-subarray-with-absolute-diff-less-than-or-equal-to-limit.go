// 1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

package leetcode

import "container/list"

func longestSubarray(nums []int, limit int) int {
	minq := list.New()
	maxq := list.New()
	left := 0
	right := 0
	for ; right < len(nums); right++ {
		for minq.Len() > 0 && minq.Back().Value.(int) > nums[right] {
			minq.Remove(minq.Back())
		}
		minq.PushBack(nums[right])
		for maxq.Len() > 0 && maxq.Back().Value.(int) < nums[right] {
			maxq.Remove(maxq.Back())
		}
		maxq.PushBack(nums[right])
		if maxq.Front().Value.(int)-minq.Front().Value.(int) > limit {
			if minq.Front().Value.(int) == nums[left] {
				minq.Remove(minq.Front())
			}
			if maxq.Front().Value.(int) == nums[left] {
				maxq.Remove(maxq.Front())
			}
			left++
		}
	}
	return right - left
}
