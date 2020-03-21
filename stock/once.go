package stock

import (
	"fmt"
)

/**
base case：
dp[-1][k][0] = dp[i][0][0] = 0
dp[-1][k][1] = dp[i][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
 */


/**
状态转移方程：
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][1-1][0] - prices[i])
			= max(dp[i-1][1][1], -prices[i]) //由于k=0,所以dp[i-1][0][0]=0
//由于k不变,消除k的影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], -prices[i])
 */
func MaxProfitOnce(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i, price := range prices {
		if i - 1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -price
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price)
		dp[i][1] = max(dp[i-1][1], - price)
		fmt.Println(dp)
	}

	return dp[len(prices)-1][0]
}

/**
状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
			= max(dp[i-1][1][1], dp[i-1][k][0] - prices[i]) //由于k=+infinity,所以k=k-1
//由于k不变,消除k的影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
 */
func MaxProfitUnlimit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i, price := range prices {
		if i - 1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -price
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - price)
	}

	return dp[len(prices)-1][0]
}

/**
状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-2][k-1][0] - prices[i])
			= max(dp[i-1][1][1], dp[i-2][k][0] - prices[i]) //由于k=+infinity,所以k=k-1
//由于k不变,消除k的影响
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
 */
func MaxProfitWithCooldown(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i, price := range prices {
		if i - 2 == -2 {
			dp[i][0] = 0
			dp[i][1] = -price
			continue
		}
		if i - 2 == -1 {
			dp[i][0] = max(dp[0][0], dp[0][1] + price)
			dp[i][1] = max(dp[0][1], -price)
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + price)
		dp[i][1] = max(dp[i-1][1], dp[i-2][0] - price)
	}

	return dp[len(prices)-1][0]
}

/**
状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
 */
func MaxProfixWithK(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	k := 2
	dp := make([][][]int, len(prices))

	for i := range dp {
		dp[i] = make([][]int, k + 1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i, price := range prices {
		for j := k; j >=1; j-=1 {
			if i - 1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -price
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + price)
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - price)
		}
	}

	return dp[len(prices)-1][k][0]
}

func MaxProfixWithAnyK(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	if k / len(prices) > 2 {
		return MaxProfitUnlimit(prices)
	}

	dp := make([][][]int, len(prices))

	for i := range dp {
		dp[i] = make([][]int, k + 1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i, price := range prices {
		for j := k; j >=1; j-=1 {
			if i - 1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -price
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + price)
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - price)
		}
	}

	return dp[len(prices)-1][k][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}