package array

import (
	"sort"
)

func CombinationSumOnce(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res = make([][]int, 0)
	path = make([]int, 0)
	dfsOnlyOnce(candidates, target, 0)
	return res
}

func dfsOnlyOnce(candidates []int, target int, begin int) {
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

		//对重复使用的剪枝
		if i > begin && candidates[i] == candidates[i - 1] {
			continue
		}

		path = append(path, candidates[i])
		//与上一题不同,不能使用重复数,所以起始点要+1
		dfsOnlyOnce(candidates, target - candidates[i], i + 1)
		path = path[:len(path) - 1]
	}
}