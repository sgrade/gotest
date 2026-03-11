// 1009. Complement of Base 10 Integer
// https://leetcode.com/problems/complement-of-base-10-integer/

package complementofbase10integer

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	bit := 1
	for i := n; i > 0; i >>= 1 {
		n = n ^ bit
		bit <<= 1
	}
	return n
}
