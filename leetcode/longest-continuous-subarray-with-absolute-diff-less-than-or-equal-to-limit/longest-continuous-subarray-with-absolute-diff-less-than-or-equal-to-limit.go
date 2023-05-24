// 1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/

package leetcode

func longestSubarray(nums []int, limit int) int {
	minQ, maxQ := []int{}, []int{}
	left := 0
	right := 0
	for ; right < len(nums); right++ {
		for len(minQ) > 0 && minQ[len(minQ)-1] > nums[right] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, nums[right])
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < nums[right] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, nums[right])
		if maxQ[0]-minQ[0] > limit {
			if minQ[0] == nums[left] {
				minQ = minQ[1:]
			}
			if maxQ[0] == nums[left] {
				maxQ = maxQ[1:]
			}
			left++
		}
	}
	return right - left
}
