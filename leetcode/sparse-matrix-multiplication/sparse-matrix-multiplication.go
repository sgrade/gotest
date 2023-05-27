// 311. Sparse Matrix Multiplication
// https://leetcode.com/problems/sparse-matrix-multiplication/

package sparsematrixmultiplication

// Based on Editorial's Approach 3: Yale Format. Optimized with ideas from Leetcode's Sample 0 ms submission.
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
	values, colIndex, rowIndex []int
}

// CSR - Compressed Sparse Row
func newCsr(matrix [][]int) *sparseMatrix {
	values, colIndex, rowIndex := []int{}, []int{}, []int{0}
	for row := range matrix {
		for col, num := range matrix[row] {
			if num != 0 {
				values = append(values, num)
				colIndex = append(colIndex, col)
			}
		}
		rowIndex = append(rowIndex, len(values))
	}
	return &sparseMatrix{values, colIndex, rowIndex}
}

// CSC - Compressed Sparse Column
func newCsc(matrix [][]int) *sparseMatrix {
	values, colIndex, rowIndex := []int{}, []int{0}, []int{}
	for col := range matrix[0] {
		for row := range matrix {
			if matrix[row][col] != 0 {
				values = append(values, matrix[row][col])
				rowIndex = append(rowIndex, row)
			}
		}
		colIndex = append(colIndex, len(values))
	}
	return &sparseMatrix{values, colIndex, rowIndex}
}
