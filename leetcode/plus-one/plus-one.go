// 66. Plus One
// https://leetcode.com/problems/plus-one/

package plusone

import "slices"

func plusOne(digits []int) []int {
	slices.Reverse(digits)
	carry := 1
	for i := range digits {
		digits[i] += carry
		if digits[i] > 9 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		digits = append(digits, 1)
	}
	slices.Reverse(digits)
	return digits
}
