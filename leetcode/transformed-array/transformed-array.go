// 3379. Transformed Array
// https://leetcode.com/problems/transformed-array/

package transformedarray

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range n {
		j := ((i+nums[i])%n + n) % n
		ans[i] = nums[j]
	}
	return ans
}
