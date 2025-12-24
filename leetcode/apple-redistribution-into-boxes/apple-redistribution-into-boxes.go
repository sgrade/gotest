// 3074. Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/

package appleredistributionintoboxes

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	apples := 0
	for _, applesInPack := range apple {
		apples += applesInPack
	}

	slices.Sort(capacity)
	slices.Reverse(capacity)

	for i, boxSize := range capacity {
		apples -= boxSize
		if apples <= 0 {
			return i + 1
		}
	}
	return len(capacity)
}
