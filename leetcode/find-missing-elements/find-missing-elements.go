// 3731. Find Missing Elements
// https://leetcode.com/problems/find-missing-elements/

package findmissingelements

import "slices"

// Sort, then append every integer that falls in a gap between consecutive values.
func findMissingElements(nums []int) []int {
	slices.Sort(nums)
	var missing []int
	for i := 0; i < len(nums)-1; i++ {
		for x := nums[i] + 1; x < nums[i+1]; x++ {
			missing = append(missing, x)
		}
	}
	return missing
}
