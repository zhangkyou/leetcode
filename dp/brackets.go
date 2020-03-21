package dp

func GenerateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
	dp := make([][]string, 0)
	dp = append(dp, []string{""}, []string{"()"})

	for i:=2; i<=n; i++ {
		tmp := make([]string, 0)
		for j:=0; j<i; j++ {
			for _, s := range dp[j] {
				for _, s1 := range dp[i-j-1] {
					cur := "(" + s + ")" + s1
					tmp = append(tmp, cur)
				}
			}
		}
		dp = append(dp, tmp)
	}

	return dp[n]
}