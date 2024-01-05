// 300. Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/

package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	sub := []int{nums[0]}
	for i := 1; i < n; i++ {
		if nums[i] > sub[len(sub)-1] {
			sub = append(sub, nums[i])
		} else {
			j := 0
			for nums[i] > sub[j] {
				j++
			}
			sub[j] = nums[i]
		}
	}
	return len(sub)
}
