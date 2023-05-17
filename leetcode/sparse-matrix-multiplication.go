// 311. Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/

package leetcode

// Based on Editorial's Approach 3: Yale Format
func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	rows, cols := len(mat1), len(mat2[0])
	ans := make([][]int, rows)
	for row := range ans {
		ans[row] = make([]int, cols)
	}

	a := newCsr(mat1)
	b := newCsc(mat2)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			mat1RowStart := a.rowIndex[row]
			mat1RowEnd := a.rowIndex[row+1]
			mat2ColStart := b.colIndex[col]
			mat2ColEnd := b.colIndex[col+1]
			for mat1RowStart < mat1RowEnd && mat2ColStart < mat2ColEnd {
				if a.colIndex[mat1RowStart] < b.rowIndex[mat2ColStart] {
					mat1RowStart++
				} else if a.colIndex[mat1RowStart] > b.rowIndex[mat2ColStart] {
					mat2ColStart++
				} else {
					ans[row][col] += a.values[mat1RowStart] * b.values[mat2ColStart]
					mat1RowStart++
					mat2ColStart++
				}

			}
		}
	}

	return ans
}

type sparseMatrix struct {
	rows                       int
	cols                       int
	values, colIndex, rowIndex []int
}

// CSR - Compressed Sparse Row
func newCsr(matrix [][]int) *sparseMatrix {
	sm := &sparseMatrix{}
	sm.rows = len(matrix)
	sm.cols = len(matrix[0])
	sm.rowIndex = append(sm.rowIndex, 0)
	for row := 0; row < sm.rows; row++ {
		for col := 0; col < sm.cols; col++ {
			if matrix[row][col] != 0 {
				sm.values = append(sm.values, matrix[row][col])
				sm.colIndex = append(sm.colIndex, col)
			}
		}
		sm.rowIndex = append(sm.rowIndex, len(sm.values))
	}
	return sm
}

// CSC - Compressed Sparse Column
func newCsc(matrix [][]int) *sparseMatrix {
	sm := &sparseMatrix{}
	sm.rows = len(matrix)
	sm.cols = len(matrix[0])
	sm.colIndex = append(sm.colIndex, 0)
	for col := 0; col < sm.cols; col++ {
		for row := 0; row < sm.rows; row++ {
			if matrix[row][col] != 0 {
				sm.values = append(sm.values, matrix[row][col])
				sm.rowIndex = append(sm.rowIndex, row)
			}
		}
		sm.colIndex = append(sm.colIndex, len(sm.values))
	}
	return sm
}
