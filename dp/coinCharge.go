package dp

import (
	"fmt"
	"math"
)

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount+1
	}

	n := len(coins)

	dp[0] = 0
	for i:=1; i<=amount; i+=1 {
		for j:=0; j<n;j+=1 {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
		if dp[i] == math.MaxInt64 {
			dp[i] = 0
		}
	}
	fmt.Println(dp)

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
