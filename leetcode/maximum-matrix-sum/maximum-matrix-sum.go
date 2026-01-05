// 2679. Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/

package maximummatrixsum

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	total, minAbs, negatives := int64(0), math.MaxInt, 0
	for _, row := range matrix {
		for _, num := range row {
			curAbs := num
			if num < 0 {
				negatives++
				total -= int64(num)
				curAbs = -num
			} else {
				total += int64(num)
			}
			minAbs = min(minAbs, curAbs)
		}
	}
	if negatives%2 != 0 {
		total -= 2 * int64(minAbs)
	}
	return total
}
