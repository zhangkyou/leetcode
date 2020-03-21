package backtrace

func Permute(nums []int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	used = make([]bool, len(nums))
	n := len(nums)

	if n == 0 {
		return res
	}
	dfs(nums, 0)

	return res
}


func dfs(nums []int, depth int) {
	if depth == len(nums) {
		add := make([]int, len(path))
		copy(add, path)
		res = append(res, add)
		return
	}

	for i:=0; i<len(nums); i+=1 {
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true

			dfs(nums, depth + 1)
			used[i] = false
			path = path[:len(path) - 1]
		}
	}
}