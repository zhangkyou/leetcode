package stdsort

func QuickSort(nums []int) []int {
	partition(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, start, end int) {
	if start >= end {
		return
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
	partition(nums, start, i-1)
	partition(nums, i+1, end)
}