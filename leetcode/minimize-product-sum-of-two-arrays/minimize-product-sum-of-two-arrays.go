// 1874. Minimize Product Sum of Two Arrays
// https://leetcode.com/problems/minimize-product-sum-of-two-arrays/

package minimizeproductsumoftwoarrays

import "sort"

func minProductSum(nums1, nums2 []int) int {
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))

	productSum := 0
	for i := range len(nums1) {
		productSum += nums1[i] * nums2[i]
	}
	return productSum
}
