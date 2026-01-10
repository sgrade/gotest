// 712. Minimum ASCII Delete Sum for Two Strings
// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/

package minimumasciideletesumfortwostrings

func minimumDeleteSum(s1 string, s2 string) int {
	n1, n2 := len(s1), len(s2)

	dp := make([][]int, n1+1)
	for i := range n1 + 1 {
		dp[i] = make([]int, n2+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + int(s1[i-1])
		}
	}
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+int(s1[i-1]),
					dp[i][j-1]+int(s2[j-1]),
				)
			}
		}
	}

	return dp[n1][n2]
}
