// 2369. Check if There is a Valid Partition For The Array
// https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/

package checkifthereisavalidpartitionforthearray

// Based on Editorial's Approach 3: Space Optimized Bottom-Up Dynamic Programming
func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]int, 3)
	dp[0] = 1

	for i := 0; i < n; i++ {
		dp_idx := i + 1
		ans := 0

		if i > 0 && nums[i] == nums[i-1] {
			ans |= dp[(dp_idx-2)%3]
		}
		if i > 1 && nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			ans |= dp[(dp_idx-3)%3]
		}
		if i > 1 && nums[i] == nums[i-1]+1 && nums[i] == nums[i-2]+2 {
			ans |= dp[(dp_idx-3)%3]
		}

		dp[dp_idx%3] = ans
	}

	return dp[n%3] != 0
}
