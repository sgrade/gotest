// 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum
// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/

package findtwononoverlappingsubarrayseachwithtargetsum

// Based on Editorial's Approach 2: Sliding Window + Dynamic Programming
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	// minLen[i] is the shortest subarray summing to target within arr[:i],
	// where n stands for "none found yet".
	minLen := make([]int, n+1)
	for i := range minLen {
		minLen[i] = n
	}
	ans := n + 1
	windowSum := 0
	left := 0

	for right, num := range arr {
		// Values are positive, so shrinking from the left lowers the sum.
		windowSum += num
		for windowSum > target {
			windowSum -= arr[left]
			left++
		}
		minLen[right+1] = minLen[right]
		if windowSum == target {
			curLen := right - left + 1
			// Pair this window with the best one ending before it.
			ans = min(ans, curLen+minLen[left])
			minLen[right+1] = min(minLen[right], curLen)
		}
	}

	if ans == n+1 {
		return -1
	}
	return ans
}
