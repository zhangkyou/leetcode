package array

func NextPermutation(nums []int)  {
	n := len(nums)
	if n == 1 {
		return
	}

	i := n - 2
	for {
		if i<0 || nums[i+1] > nums[i] {
			break
		}
		i -= 1
	}
	if i >= 0 {
		j := n - 1
		for {
			if j < 0 || nums[j] > nums[i] {
				break
			}
			j -= 1
		}
		swap(nums, i, j)
	}
	reverse(nums, i + 1)

	return
}

func reverse(nums []int, start int) {
	n := len(nums)

	i := start
	j := n - 1
	for {
		if i >= j {
			break
		}

		swap(nums, i, j)
		i += 1
		j -= 1
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}