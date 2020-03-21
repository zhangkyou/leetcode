package array

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0
	preMin := prices[0]

	for _, num := range prices {
		cur := num - preMin
		if cur > res {
			res = cur
		}
		if num < preMin {
			preMin = num
		}
	}

	return res
}

func MaxProfitMulti(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	res := 0
	for i:=1; i<len(prices); i+=1 {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}