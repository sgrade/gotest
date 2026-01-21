// 3315. Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/

package constructtheminimumbitwisearrayii

func minBitwiseArray(nums []int) []int {
	for i, num := range nums {
		cur := -1
		bit := 1
		for (num & bit) != 0 {
			cur = num - bit
			bit <<= 1
		}
		nums[i] = cur
	}
	return nums
}
