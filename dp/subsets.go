package dp

import (
	"fmt"
)

func Subsets(nums []int) [][]int {
	n := len(nums)

	dp := make([][]int, 0)

	dp = append(dp, []int{})
	for i:=0; i<n; i++ {
		for _, tmp := range dp {
			cur := append(tmp, nums[i])
			add := make([]int, len(cur))
			copy(add, cur)
			dp = append(dp, add)
		}
	}
	return dp
}

func appendTrick() {
	dp := make([]int, 0)
	fmt.Printf("%v, %p, %d, %d\n", dp, dp, len(dp), cap(dp))

	dp1 := append(dp, 1)
	fmt.Printf("%v, %p, %d, %d\n", dp1, dp1, len(dp1), cap(dp1))

	dp2 := append(dp1, 2)
	fmt.Printf("%v, %p, %d, %d\n", dp2, dp2, len(dp2), cap(dp2))

	dp3 := append(dp2, 3)
	fmt.Printf("%v, %p, %d, %d\n", dp3, dp3, len(dp3), cap(dp3))

	dp4 := append(dp3, 4)
	fmt.Printf("dp4:%v, %p, %d, %d\n", dp4, dp4, len(dp4), cap(dp4))

	dp5 := append(dp3, 5)
	fmt.Printf("dp5:%v, %p, %d, %d\n", dp5, dp5, len(dp5), cap(dp5))

	fmt.Printf("dp4:%v, %p, %d, %d\n", dp4, dp4, len(dp4), cap(dp4))
}