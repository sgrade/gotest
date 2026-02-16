// 190. Reverse Bits
// https://leetcode.com/problems/reverse-bits/

package reversebits

func reverseBits(n int) int {
	ans, power := 0, 31
	for n != 0 {
		ans += (n & 1) << power
		n = n >> 1
		power -= 1
	}
	return ans
}
