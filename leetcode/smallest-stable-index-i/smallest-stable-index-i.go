// 3903. Smallest Stable Index I
// https://leetcode.com/problems/smallest-stable-index-i/

package smalleststableindexi

func firstStableIndex(nums []int, k int) int {
	n := len(nums)

	// suffixMin[i] is the minimum of nums[i:].
	suffixMin := make([]int, n)
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(nums[i], suffixMin[i+1])
	}

	// Sweep left to right; prefixMax is the maximum of nums[:i+1].
	prefixMax := nums[0]
	for i, num := range nums {
		prefixMax = max(prefixMax, num)
		if prefixMax-suffixMin[i] <= k {
			return i
		}
	}
	return -1
}
