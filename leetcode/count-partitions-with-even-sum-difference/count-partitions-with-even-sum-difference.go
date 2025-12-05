// 3432. Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/

package countpartitionswithevensumdifference

func countPartitions(nums []int) int {
	sum := 0
	for i := range len(nums) {
		sum += nums[i]
	}

	left, ans := 0, 0
	for i := range len(nums) - 1 {
		left += nums[i]
		right := sum - left
		if (right-left)%2 == 0 {
			ans += 1
		}
	}
	return ans
}
