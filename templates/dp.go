package templates

// DP - 1D Array
func dp1D(n int) {
	dp := make([]int, n+1)
	dp[0] = 0 // base case
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // recurrence
	}
}

// DP - 2D Grid
func dp2D(m, n int) {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// base cases
	dp[0][0] = 0
	// fill dp table
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// dp[i][j] = ...
		}
	}
}

// DP - Min/Max Pattern
func minMaxDP(arr []int) {
	n := len(arr)
	dp := make([]int, n)
	dp[0] = arr[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+arr[i], arr[i]) // or min
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

