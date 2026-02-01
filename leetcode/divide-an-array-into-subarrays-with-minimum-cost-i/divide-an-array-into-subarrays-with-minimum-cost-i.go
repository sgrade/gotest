// 3010. Divide an Array Into Subarrays With Minimum Cost I
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/

package divideanarrayintosubarrayswithminimumcosti

import "slices"

func minimumCost(nums []int) int {
	cost := nums[0]
	mins := []int{nums[1], nums[2]}
	slices.Sort(mins)
	for i := 3; i < len(nums); i++ {
		if nums[i] < mins[1] {
			mins[1] = nums[i]
			slices.Sort(mins)
		}
	}
	cost += mins[0] + mins[1]
	return cost
}
