// 311. Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/

package leetcode

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	rows, kk, cols := len(mat1), len(mat1[0]), len(mat2[0])
	ans := make([][]int, rows)
	for row := range ans {
		ans[row] = make([]int, cols)
	}
	for row := 0; row < rows; row++ {
		for k := 0; k < kk; k++ {
			if mat1[row][k] != 0 {
				for col := 0; col < cols; col++ {
					ans[row][col] += mat1[row][k] * mat2[k][col]
				}
			}
		}
	}
	return ans
}
