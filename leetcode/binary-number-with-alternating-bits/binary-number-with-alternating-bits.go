package binarynumberwithalternatingbits

func hasAlternatingBits(n int) bool {
	prev := n & 1
	x := 2
	for x <= n {
		cur := n & x
		if (cur != 0) == (prev != 0) {
			return false
		}
		prev = cur
		x = x << 1
	}
	return true
}
