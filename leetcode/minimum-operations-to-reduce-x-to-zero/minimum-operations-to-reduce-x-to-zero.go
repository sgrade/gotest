// 1658. Minimum Operations to Reduce X to Zero
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

package minimumoperationstoreducextozero

func minOperations(nums []int, x int) int {
	// Removing a prefix and a suffix summing to x keeps a middle subarray
	// summing to sum(nums) - x, so minimizing the removals means finding
	// the longest such subarray.
	sum := 0
	for _, num := range nums {
		sum += num
	}
	target := sum - x
	if target < 0 {
		return -1
	}

	longest, windowSum, left := -1, 0, 0
	for right, num := range nums {
		windowSum += num
		// Values are positive, so shrinking from the left is enough to
		// bring an oversized window back to at most target.
		for windowSum > target {
			windowSum -= nums[left]
			left++
		}
		if windowSum == target {
			longest = max(longest, right-left+1)
		}
	}
	if longest < 0 {
		return -1
	}
	return len(nums) - longest
}
