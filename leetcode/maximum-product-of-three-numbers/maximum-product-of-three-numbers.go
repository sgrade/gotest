// 628. Maximum Product of Three Numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/

package maximumproductofthreenumbers

import "sort"

func maximumProduct(nums []int) int {
	mx := []int{-1000, -1000, -1000}
	mn := []int{1000, 1000}
	for _, num := range nums {
		if num > mx[0] {
			mx = append(mx, num)
			sort.Ints(mx)
			mx = mx[1:]
		}
		if num < mn[1] {
			mn = append(mn, num)
			sort.Ints(mn)
			mn = mn[:2]
		}
	}
	return max(mx[2]*mx[1]*mx[0], mx[2]*mn[0]*mn[1])
}
