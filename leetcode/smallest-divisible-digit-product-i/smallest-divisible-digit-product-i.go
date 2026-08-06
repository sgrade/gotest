// 3345. Smallest Divisible Digit Product I
// https://leetcode.com/problems/smallest-divisible-digit-product-i/

package smallestdivisibledigitproducti

func smallestNumber(n int, t int) int {
	for x := n; x < n+10; x++ {
		product := 1
		tmp := x
		for tmp > 0 {
			product *= tmp % 10
			tmp /= 10
		}
		if product%t == 0 {
			return x
		}
	}
	return -1
}
