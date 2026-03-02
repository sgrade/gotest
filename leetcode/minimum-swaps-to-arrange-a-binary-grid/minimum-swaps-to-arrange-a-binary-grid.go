// 1536. Minimum Swaps to Arrange a Binary Grid
// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/

package minimumswapstoarrangeabinarygrid

// Based on Editorial's Approach: Greedy
func minSwaps(grid [][]int) int {
	n := len(grid)
	lastOne := make([]int, n)
	for i := range lastOne {
		lastOne[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				lastOne[i] = j
				break
			}
		}
	}

	swaps := 0
	for i := 0; i < n; i++ {
		k := -1
		for j := i; j < n; j++ {
			if lastOne[j] <= i {
				swaps += j - i
				k = j
				break
			}
		}

		if k != -1 {
			for j := k; j > i; j-- {
				lastOne[j], lastOne[j-1] = lastOne[j-1], lastOne[j]
			}
		} else {
			return -1
		}
	}

	return swaps
}
