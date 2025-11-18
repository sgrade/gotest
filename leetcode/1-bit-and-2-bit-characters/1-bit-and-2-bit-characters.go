package bitand2bitcharacters

// 717. 1-bit and 2-bit Characters
// https://leetcode.com/problems/1-bit-and-2-bit-characters/

func isOneBitCharacter(bits []int) bool {
	ones := 0
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 1 {
			ones++
		} else {
			break
		}
	}
	return ones%2 == 0
}
