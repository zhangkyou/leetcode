package dp

func MaxPackageValue(w int, weights, values[]int) int {
	n := len(weights)
	dp := make([][]int, n + 1)
	for i := 0;i < n+1;i += 1 {
		dp[i] = make([]int, w + 1)
	}

	for i := 1; i < n+1; i+=1 {
		for j:=1; j<w + 1; j+=1 {
			if j - weights[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}

	return dp[n][w]
}