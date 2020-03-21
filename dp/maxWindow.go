package dp

import (
	"math"

	"Goproject/leetcode/queue"
)

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	dp := make([]int, len(nums) - k + 1)

	for k := range dp {
		dp[k] = math.MinInt8
	}

	for i:=0; i< len(nums); i+=1 {

		for j:=1; j<=k; j+=1 {
			if i - k + j >= 0 && i - k + j < len(dp) {
				dp[i - k + j] = max(dp[i - k + j], nums[i])
			}
		}
	}

	return dp
}

func MaxSlidingWindowOptimiz(nums []int, k int) []int {
	res := make([]int, 0)

	q := queue.Constructor()

	for i:=0; i<len(nums); i+=1 {
		q.Push_back(nums[i])
		if i >= k - 1 {
			res = append(res, q.Max_value())
			q.Pop_front()
		}
	}

	return res
}

