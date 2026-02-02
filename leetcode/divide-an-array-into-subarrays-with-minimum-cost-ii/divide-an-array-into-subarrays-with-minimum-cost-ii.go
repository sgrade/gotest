package divideanarrayintosubarrayswithminimumcostii

import "github.com/sgrade/gotest/templates/structure"

// Based on the editorial's Approach: Ordered Set
func minimumCost(nums []int, k, dist int) int64 {
	n := len(nums)
	// Track k-2 smallest elements (we always include nums[0] and one other element)
	tracker := structure.NewKSmallestTracker(k - 2)

	// Initialize tracker with first k-1 elements (excluding nums[0])
	for i := 1; i < k-1; i++ {
		tracker.Add(nums[i])
	}

	// Start with cost of nums[0] + (k-2 smallest from nums[1:k-1]) + nums[k-1]
	ans := int64(nums[0]) + tracker.Sum() + int64(nums[k-1])

	// Sliding window: maintain k-2 smallest elements in valid range
	for i := k; i < n; i++ {
		// Remove element that falls outside the distance constraint
		j := i - dist - 1
		if j > 0 {
			tracker.Remove(nums[j])
		}
		// Add the element at position i-1 to the window
		tracker.Add(nums[i-1])
		// Calculate cost: nums[0] + (k-2 smallest) + nums[i]
		if current := int64(nums[0]) + tracker.Sum() + int64(nums[i]); current < ans {
			ans = current
		}
	}

	return ans
}
