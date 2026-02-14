// 799. Champagne Tower
// https://leetcode.com/problems/champagne-tower/

package champagnetower

// Based on Editorial's Approach #1: Simulation
func champagneTower(poured int, query_row int, query_glass int) float64 {
	glasses := make([][]float64, 102)
	for i := range glasses {
		glasses[i] = make([]float64, 102)
	}
	glasses[0][0] = float64(poured)
	for r := 0; r < query_row; r++ {
		for c := 0; c <= r; c++ {
			champagneToFall := (glasses[r][c] - 1.0) / 2.0
			if champagneToFall > 0 {
				glasses[r+1][c] += champagneToFall
				glasses[r+1][c+1] += champagneToFall
			}
		}
	}
	return min(1, glasses[query_row][query_glass])
}
