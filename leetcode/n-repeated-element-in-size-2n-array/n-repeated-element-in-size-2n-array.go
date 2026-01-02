// 961. N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

package nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) int {
	for diff := 1; diff < 4; diff++ {
		for i := 0; i < len(nums)-diff; i++ {
			if nums[i] == nums[i+diff] {
				return nums[i]
			}
		}
	}
	return -1
}
