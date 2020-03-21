package dp

import (
	"fmt"
)

var dp [][]int

func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i:=1; i<=m; i+=1 {
		dp[i][0] = i
	}

	for j:=1; j<=n; j+=1 {
		dp[0][j] = j
	}

	for i:=1; i<=m; i+=1 {
		for j:=1; j<=n; j+=1 {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minThree(dp[i-1][j] + 1, dp[i][j-1] + 1, dp[i-1][j-1] + 1)
			}
		}
	}
	fmt.Println(dp)

	return dp[m][n]
}

func minThree(x, y, z int) int {
	return min(x, min(y, z))
}

