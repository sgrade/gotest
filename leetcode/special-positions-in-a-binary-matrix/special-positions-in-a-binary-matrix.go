// 1582. Special Positions in a Binary Matrix
// https://leetcode.com/problems/special-positions-in-a-binary-matrix/

package specialpositionsinabinarymatrix

func numSpecial(mat [][]int) int {
	candidateCols := make([]int, len(mat[0]))

	for r := range mat {
		for c := range mat[0] {
			if mat[r][c] == 1 {
				candidateCols[c]++
			}
		}
	}

	specialPositions := 0
	for c, count := range candidateCols {
		if count != 1 {
			continue
		}
		// Find the row that has the 1 in this column
		rowWithOne := -1
		for r := range mat {
			if mat[r][c] == 1 {
				rowWithOne = r
				break
			}
		}
		// Check that row also has exactly one 1
		rowOnes := 0
		for col := range mat[0] {
			if mat[rowWithOne][col] == 1 {
				rowOnes++
			}
		}
		if rowOnes == 1 {
			specialPositions++
		}
	}
	return specialPositions
}
