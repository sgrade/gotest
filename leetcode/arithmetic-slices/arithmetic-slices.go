// 413. Arithmetic Slices
// https://leetcode.com/problems/arithmetic-slices/

package arithmeticslices

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := 0
	diff := nums[1] - nums[0]
	cur_len := 2
	for i := 2; i < n; i++ {
		cur_diff := nums[i] - nums[i-1]
		if cur_diff == diff {
			cur_len++
			ans += cur_len - 2
		} else {
			cur_len = 2
		}
		diff = cur_diff
	}
	return ans
}
