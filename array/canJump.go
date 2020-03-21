package array

import (
	"Goproject/leetcode/util"
)

func CanJump(nums []int) bool {
	k := 0

	for i:=0; i < len(nums); i+=1 {
		if i > k {
			return false
		}

		k = util.Max(k, nums[i] + i)
	}

	return true
}
