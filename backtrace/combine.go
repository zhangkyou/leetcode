package backtrace

func Combine(n int, k int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)

	dfsCombine(n, 1, k)
	return res
}

func dfsCombine(n int, begin int, k int) {
	if len(path) == k {
		add := make([]int, len(path))
		copy(add, path)
		res = append(res, add)
		return
	}

	for i:=begin; i<=n; i+=1 {
		path = append(path, i)

		dfsCombine(n, i + 1, k)
		path = path[:len(path) - 1]
	}
}