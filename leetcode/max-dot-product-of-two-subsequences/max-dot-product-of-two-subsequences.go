// 1458. Max Dot Product of Two Subsequences
// https://leetcode.com/problems/max-dot-product-of-two-subsequences/

package maxdotproductoftwosubsequences

import "slices"

// Based on Editorial's Approach 1: Top-Down Dynamic Programming
func maxDotProduct(nums1 []int, nums2 []int) int {
	min1, max1 := slices.Min(nums1), slices.Max(nums1)
	min2, max2 := slices.Min(nums2), slices.Max(nums2)

	// Early return for edge cases
	if max1 < 0 && min2 > 0 {
		return max1 * min2
	}
	if max2 < 0 && min1 > 0 {
		return max2 * min1
	}
	// If one array is all zeros, result is 0
	if max1 == 0 && min1 == 0 {
		return 0
	}
	if max2 == 0 && min2 == 0 {
		return 0
	}

	const SENTINEL = -1_000_000_000 // Use int constant, not float
	memo := make([][]int, len(nums1))
	for i := range nums1 {
		memo[i] = make([]int, len(nums2))
		for j := range nums2 {
			memo[i][j] = SENTINEL
		}
	}

	var getDotProduct func(idx1, idx2 int) int
	getDotProduct = func(idx1, idx2 int) int {
		if idx1 == len(nums1) || idx2 == len(nums2) {
			return 0 // Base case: no more pairs to add
		}

		if memo[idx1][idx2] != SENTINEL {
			return memo[idx1][idx2]
		}

		// Take current pair and optionally continue (max(0, ...) allows early termination)
		takeBoth := nums1[idx1] * nums2[idx2]
		moveBoth := takeBoth + max(0, getDotProduct(idx1+1, idx2+1))

		// Skip current element in nums1, optionally continue
		moveIdx1 := max(0, getDotProduct(idx1+1, idx2))

		// Skip current element in nums2, optionally continue
		moveIdx2 := max(0, getDotProduct(idx1, idx2+1))

		memo[idx1][idx2] = max(moveBoth, max(moveIdx1, moveIdx2))
		return memo[idx1][idx2]
	}

	return getDotProduct(0, 0)
}
