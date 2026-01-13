// 3453. Separate Squares I
// https://leetcode.com/problems/separate-squares-i/

package separatesquaresi

// Binary search to find the horizontal line that divides total area in half.
// Time: O(n log maxY), Space: O(1)
func separateSquares(squares [][]int) float64 {
	maxY, totalArea := 0.0, 0.0
	for _, sq := range squares {
		y, side := float64(sq[1]), float64(sq[2])
		totalArea += side * side
		maxY = max(maxY, y+side)
	}

	hasHalfAreaBelow := func(targetY float64) bool {
		areaBelowTarget := 0.0
		for _, sq := range squares {
			y, side := float64(sq[1]), float64(sq[2])
			if y < targetY {
				areaBelowTarget += side * min(targetY-y, side)
			}
		}
		return areaBelowTarget >= totalArea/2
	}

	lo, hi := 0.0, maxY
	precision := 1e-5
	for hi-lo > precision {
		mid := lo + (hi-lo)/2
		if hasHalfAreaBelow(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}
