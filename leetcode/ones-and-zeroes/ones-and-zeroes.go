package onesandzeroes

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeroes := strings.Count(s, "0")
		ones := len(s) - zeroes

		for z := m; z >= zeroes; z-- {
			for o := n; o >= ones; o-- {
				dp[z][o] = max(dp[z][o], 1+dp[z-zeroes][o-ones])
			}
		}
	}

	return dp[m][n]
}
