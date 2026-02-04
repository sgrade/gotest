// 3640. Trionic Array II
// https://leetcode.com/problems/trionic-array-ii/

package trionicarrayii

import "math"

// Based on Editorial's Approach: Grouped Loop
func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	var p, q, j int
	var maxSum, sum, curAns int64
	ans := int64(math.MinInt64)

	for i := 0; i < n; i++ {
		j = i + 1
		curAns = 0

		// first segment
		for ; j < n && nums[j-1] < nums[j]; j++ {
		}
		p = j - 1
		if p == i {
			continue
		}

		// second segment
		curAns += int64(nums[p] + nums[p-1])
		for ; j < n && nums[j-1] > nums[j]; j++ {
			curAns += int64(nums[j])
		}
		q = j - 1
		if q == p || q == n-1 || (j < n && nums[j] <= nums[q]) {
			i = q
			continue
		}

		// third segment
		curAns += int64(nums[q+1])
		maxSum = 0
		sum = 0
		for k := q + 2; k < n && nums[k] > nums[k-1]; k++ {
			sum += int64(nums[k])
			maxSum = max(maxSum, sum)
		}
		curAns += maxSum

		// max sum in first segment
		maxSum = 0
		sum = 0
		for k := p - 2; k >= i; k-- {
			sum += int64(nums[k])
			maxSum = max(maxSum, sum)
		}
		curAns += maxSum

		ans = max(ans, curAns)
		i = q - 1
	}
	return ans
}
