package array

import (
	"sort"
)
var res [][]int
var path []int

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res = make([][]int, 0)
	path = make([]int, 0)
	dfs(candidates, target, 0)
	return res
}

func dfs(candidates []int, target int, begin int) {
	if target == 0 {
		add := make([]int, len(path))
		copy(add, path)
		res = append(res, add)
		return
	}
	n := len(candidates)

	for i := begin; i < n; i+=1 {
		if target - candidates[i] < 0 {
			break
		}

		path = append(path, candidates[i])
		dfs(candidates, target - candidates[i], i)
		path = path[:len(path) - 1]
	}
}