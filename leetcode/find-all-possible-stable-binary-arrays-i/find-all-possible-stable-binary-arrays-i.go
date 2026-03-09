// 3129. Find All Possible Stable Binary Arrays I
// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-i/

package findallpossiblestablebinaryarraysi

// Based on Editorial's Approach: Dynamic Programming
func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007

	// dp[i][j][v] = ways to arrange i zeros and j ones such that the last
	// element is v and no consecutive run exceeds limit.
	dp := make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	// Base cases: an array of all zeros (or all ones) up to length limit.
	for i := 1; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 1; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}

	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			// Append a 0: extend any (i-1, j) array.
			dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			// Subtract arrangements that would create a run of limit+1 zeros.
			// Those are the ones that had exactly limit zeros already and ended
			// with a 1 before that run, i.e. dp[i-limit-1][j][1].
			if i > limit {
				dp[i][j][0] -= dp[i-limit-1][j][1]
			}
			dp[i][j][0] = (dp[i][j][0]%mod + mod) % mod

			// Append a 1: symmetric logic along the j axis.
			dp[i][j][1] = dp[i][j-1][1] + dp[i][j-1][0]
			if j > limit {
				dp[i][j][1] -= dp[i][j-limit-1][0]
			}
			dp[i][j][1] = (dp[i][j][1]%mod + mod) % mod
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
