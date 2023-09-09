// 118. Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/

package pascalstriangle

func generate(numRows int) [][]int {
	ans := [][]int{}
	ans = append(ans, []int{1})
	for i := 1; i < numRows; i++ {
		currentRow := make([]int, i+1)
		currentRow[0] = 1
		j := 1
		for ; j < i; j++ {
			currentRow[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		currentRow[j] = 1
		ans = append(ans, currentRow)
	}
	return ans
}
