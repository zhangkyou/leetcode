package dp

type pair struct {
	Fir int
	Sec int
}

func StoneGame(piles []int) bool {
	dp := make([][]*pair, len(piles))
	for i := range dp {
		dp[i] = make([]*pair, len(piles))
	}

	for i:=0; i<len(piles); i++ {
		for j:=0; j<len(piles); j++ {
			dp[i][j] = &pair{0,0}
		}
	}

	for i:=0; i<len(piles); i++ {
		dp[i][i].Fir = piles[i]
		dp[i][i].Sec = 0
	}

	for k:=1; k<len(piles); k++ {
		for i:=0; i<len(piles); i++ {
			j := i + 1
			if j > len(piles) - 1 {
				continue
			}
			left := piles[i] + dp[i+1][j].Sec
			right := piles[j] + dp[i][j-1].Sec
			if left > right {
				dp[i][j].Fir = left
				dp[i][j].Sec = dp[i+1][j].Fir
			} else {
				dp[i][j].Fir = right
				dp[i][j].Sec = dp[i][j-1].Fir
			}
		}
	}

	//for l := 2; l <= len(piles); l++ {
	//	for i := 0; i < len(piles) - 1; i++ {
	//		j := l + i - 1
	//		left := piles[i] + dp[i+1][j].Sec
	//		right := piles[j] + dp[i][j-1].Sec
	//		if left > right {
	//			dp[i][j].Fir = left
	//			dp[i][j].Sec = dp[i+1][j].Fir
	//		} else {
	//			dp[i][j].Fir = right
	//			dp[i][j].Sec = dp[i][j-1].Fir
	//		}
	//	}
	//}
	return dp[len(piles)-1][len(piles)-1].Fir - dp[len(piles)-1][len(piles)-1].Sec > 0
}
