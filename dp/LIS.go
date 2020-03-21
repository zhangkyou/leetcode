package dp

func LengthOfLIS(nums []int) int {
	n := len(nums)

	res :=0
	dp := make([]int, n)
	for k := range dp {
		dp[k] = 1
	}

	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func NumberOfLIS(nums []int) int {
	n := len(nums)

	res :=0
	dp := make([]int, n)
	count := make([]int, n)
	for k := range dp {
		dp[k] = 1
	}

	for k := range count {
		count[k] = 1
	}

	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			if nums[j] < nums[i] {
				if dp[j] + 1 > dp[i] {
					count[i] = count[j]
				} else if dp[j] + 1 == dp[i] {
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}

	ansi := 0
	for k, value := range dp {
		if value == res {
			ansi += count[k]
		}
	}

	return ansi
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}