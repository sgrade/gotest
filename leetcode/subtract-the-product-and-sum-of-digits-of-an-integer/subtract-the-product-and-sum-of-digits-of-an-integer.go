// 1281. Subtract the Product and Sum of Digits of an Integer
// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

package subtracttheproductandsumofdigitsofaninteger

import "strconv"

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	s := strconv.Itoa(n)
	for _, ch := range s {
		digit := int(ch - '0')
		sum += digit
		product *= digit
	}
	return product - sum
}
