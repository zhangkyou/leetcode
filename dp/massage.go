package dp

func Massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[1]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]

	res := dp[0]
	for i:=1; i < len(nums); i+=1 {
		for j:=0; j < i; j+=1 {
			if i > j + 1 {
				dp[i] = max(dp[i], dp[j] + nums[i])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
