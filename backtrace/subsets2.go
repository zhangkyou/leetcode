package backtrace

import (
	"sort"
)

func Backtrace2(nums []int) [][]int {
	res = make([][]int, 0)
	temp = make([]int, 0)

	sort.Ints(nums)
	traceUnique(nums, 0)
	return res
}

func traceUnique(nums []int, begin int) {
	add := make([]int, len(temp))
	copy(add, temp)
	res = append(res, add)

	for i:=begin;i < len(nums); i+=1 {
		if i>begin && nums[i] == nums[i - 1] {
			continue
		}
		temp = append(temp, nums[i])

		traceUnique(nums, i+1)
		temp = temp[:len(temp) - 1]
	}
}

