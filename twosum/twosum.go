package twosum

func Sum(nums []int, target int) []int {
	tmp := make(map[int]int, 0)
	result := make([]int, 0)
	for k, v := range nums {
		tmp[target-v] = k
	}

	for k, v := range nums {
		value, ok := tmp[v]
		if ok && k != value {
			result = append(result, k, value)
			break
		}
	}

	return result
}

