package minimumremovalstobalancearray

import "sort"

func minRemoval(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)
	ans := n
	right := 0
	for left := 0; left < n; left++ {
		for right < n && int64(nums[right]) <= int64(nums[left])*int64(k) {
			right++
		}
		cur_ans := n - (right - left)
		ans = min(ans, cur_ans)
	}
	return ans
}
