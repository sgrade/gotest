// 598. Range Addition II
// https://leetcode.com/problems/range-addition-ii/

package rangeadditionii

func maxCount(m int, n int, ops [][]int) int {
	minRow, minCol := m, n
	for _, operation := range ops {
		minRow = min(minRow, operation[0])
		minCol = min(minCol, operation[1])
	}
	return minRow * minCol
}
