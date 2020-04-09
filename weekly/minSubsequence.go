package weekly

import "sort"

func MinSubsequence(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)

	i := 0
	j := n - 1
	leftSum := nums[i]
	rightSum := nums[j]
	for {
		if i >= j {
			break
		}

		if leftSum < rightSum {
			i += 1
			leftSum += nums[i]
		} else {
			j -=1
			rightSum += nums[j]
		}
	}

	res := make([]int, 0)
	for i := n - 1; i >= j; i-=1 {
		res = append(res, nums[i])
	}
	return res
}
