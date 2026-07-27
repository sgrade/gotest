// 1464. Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/

package maximumproductoftwoelementsinanarray

func maxProduct(nums []int) int {
	mx1, mx2 := 0, 0
	for _, num := range nums {
		if num > mx1 {
			mx1 = num
		}
		if mx1 > mx2 {
			mx1, mx2 = mx2, mx1
		}
	}
	return (mx1 - 1) * (mx2 - 1)
}
