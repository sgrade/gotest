// 3345. Smallest Divisible Digit Product I
// https://leetcode.com/problems/smallest-divisible-digit-product-i/

package smallestdivisibledigitproducti

// An answer always exists in [n, n+9]: one of those numbers ends in 0,
// so its digit product is 0 and divisible by any t.
func smallestNumber(n int, t int) int {
	for ; digitProduct(n)%t != 0; n++ {
	}
	return n
}

func digitProduct(x int) int {
	p := 1
	for ; x > 0; x /= 10 {
		p *= x % 10
	}
	return p
}
