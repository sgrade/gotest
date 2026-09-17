// 1762. Buildings With an Ocean View
// https://leetcode.com/problems/buildings-with-an-ocean-view/

package buildingswithanoceanview

import "slices"

func findBuildings(heights []int) []int {
	var ans []int
	maxHeight := 0
	for i, height := range slices.Backward(heights) {
		if height > maxHeight {
			ans = append(ans, i)
			maxHeight = height
		}
	}
	slices.Reverse(ans)
	return ans
}
