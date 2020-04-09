package weekly

var path []int
var res int

func NumTeams(rating []int) int {
	path = make([]int, 0)
	res = 0

	dfs(rating, 0, true)
	dfs(rating, 0, false)
	return res
}

func dfs(rating []int, start int, ascend bool) {
	if len(path) == 3 {
		res += 1
		return
	}
	n := len(rating)
	for i:=start; i < n; i+=1 {
		if ascend == true && len(path) > 0 && rating[i] < path[len(path)-1] {
			continue
		} else if ascend == false && len(path) > 0 && rating[i] > path[len(path)-1] {
			continue
		}

		path = append(path, rating[i])
		dfs(rating, i+1, ascend)
		path = path[:len(path)-1]
	}
	return
}
