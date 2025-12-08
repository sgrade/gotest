package countsquaresumtriples

import "math"

// Based on Editorial's Approach: Enumeration
func countTriples(n int) int {
	ans := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			// We use +1 to avoid floating point inaccuracies and take into account that the gap between consequitive perfect squares >= 1
			c := int(math.Sqrt(float64(a*a + b*b + 1)))
			if c <= n && c*c == a*a+b*b {
				ans++
			}
		}
	}
	return ans
}
