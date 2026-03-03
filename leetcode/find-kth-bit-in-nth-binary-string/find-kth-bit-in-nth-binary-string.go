// 1545. Find Kth Bit in Nth Binary String
// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/

package findkthbitinnthbinarystring

// Based on Editorial's Approach 3: Iterative Divide and Conquer
func findKthBit(n int, k int) byte {
	invertCount := 0
	len := (1 << n) - 1

	for k > 1 {
		// middle
		if k == len/2+1 {
			if invertCount%2 == 0 {
				return '1'
			}
			return '0'
		}

		if k > len/2 {
			k = len + 1 - k // mirror
			invertCount++
		}

		len /= 2
	}

	if invertCount%2 == 0 {
		return '0'
	}
	return '1'
}
