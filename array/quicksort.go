package array

func QuickSort(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	if len(nums) < k {
		return nums
	}

	partitionArray(nums, 0, len(nums) - 1, k)

	res := make([]int, 0)
	for i := 0; i < k; i += 1{
		res = append(res, nums[i])
	}

	return res
}

func partitionArray(nums []int, start, end, k int) {
	m := partition(nums, start, end)
	if k == m {
		return
	} else if k < m {
		partitionArray(nums, start, m - 1, k)
	} else {
		partitionArray(nums, m + 1, end, k)
	}
}

func partition(nums []int, start, end int) int {
	if start >= end {
		return start
	}

	pre := nums[start]

	i := start
	j := end
	for {
		if i >= j {
			break
		}

		for {
			if j <= i {
				break
			}

			if nums[j] < pre {
				break
			}

			j--
		}
		nums[i] = nums[j]

		for {
			if i >= j {
				break
			}

			if nums[i] > pre {
				break
			}

			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pre
	return i
}
