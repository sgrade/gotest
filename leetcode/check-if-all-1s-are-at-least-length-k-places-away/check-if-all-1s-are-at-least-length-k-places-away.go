package checkifall1sareatleastlengthkplacesaway

// 1437. Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

import "slices"

func kLengthApart(nums []int, k int) bool {
	prev := slices.Index(nums, 1)
	for cur := prev + 1; cur < len(nums); cur++ {
		if nums[cur] == 1 {
			dist := cur - prev - 1
			if dist < k {
				return false
			}
			prev = cur
		}
	}
	return true
}
