package backtrace

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(nums))
	n := len(nums)
	sort.Ints(nums)
	if n == 0 {
		return res
	}
	dfsUnique(nums, 0)

	return res
}


func dfsUnique(nums []int, depth int) {
	if depth == len(nums) {
		add := make([]int, len(path))
		copy(add, path)
		res = append(res, add)
		return
	}

	for i:=0; i<len(nums); i+=1 {
		if !used[i] {
			if i>0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true

			dfsUnique(nums, depth + 1)
			used[i] = false
			path = path[:len(path) - 1]
		}
	}
}