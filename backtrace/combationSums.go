package backtrace

import (
	"fmt"
	"math"
	"sort"
)

var resCharge int
var pathLen int

func CoinChange(coins []int, amount int) int {
	pathLen = 0
	resCharge = math.MaxInt64

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	fmt.Println(coins)

	dfsCharge(coins, amount)
	if resCharge == math.MaxInt64 && amount != 0 {
		return -1
	}
	return resCharge
}

//dfs思想会很耗时
func dfsCharge(coins []int, target int) {
	if target == 0 {
		if pathLen < resCharge {
			resCharge = pathLen
		}
		return
	}

	if target < 0 {
		return
	}

	if pathLen > resCharge {
		return
	}

	for i:=0; i<len(coins); i+=1 {
		pathLen += 1

		cur := target - coins[i]
		dfsCharge(coins, cur)
		pathLen -= 1
	}
}