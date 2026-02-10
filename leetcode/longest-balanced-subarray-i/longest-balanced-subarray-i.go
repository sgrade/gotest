// 3719. Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/

package longestbalancedsubarrayi

func longestBalanced(nums []int) int {
	maxLen := 0
	for left := 0; left < len(nums); left++ {
		odd := make(map[int]int)
		even := make(map[int]int)
		for right := left; right < len(nums); right++ {
			if nums[right]&1 == 1 {
				odd[nums[right]]++
			} else {
				even[nums[right]]++
			}
			if len(odd) == len(even) {
				maxLen = max(maxLen, right-left+1)
			}
		}
	}
	return maxLen
}
