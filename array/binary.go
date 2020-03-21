package array

func Search(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	i := 0
	j := n - 1

	if target < nums[i] && target > nums[j] {
		return -1
	}

	if target >= nums[i] {
		for {
			if i > j {
				break
			}

			m := i + (j - i)/2
			if nums[m] == target {
				return m
			}
			if nums[m] > target {
				return BinarySearchOptimize(nums, i, m, target)
			} else {
				if nums[m] >= nums[i] {
					i = m + 1
				} else {
					j = m - 1
				}
			}
		}
	} else {
		for {
			if i > j {
				break
			}

			m := i + (j - i)/2
			if nums[m] == target {
				return m
			}
			if nums[m] < target {
				return BinarySearchOptimize(nums, m, j, target)
			} else {
				if nums[m] > nums[j] {
					i = m + 1
				} else {
					j = m - 1
				}
			}
		}
	}
	return -1
}


func BinarySearchOptimize(nums []int, start, end, target int) int {
	i := start
	j := end

	for {
		if i > j {
			break
		}

		m := i + (j - i)/2
		if target == nums[m] {
			return m
		}
		if target < nums[m] {
			j = m - 1
		} else if target >nums[m] {
			i = m + 1
		}
	}

	return -1
}

func BinarySearch(nums []int, target int) int {
	n := len(nums)

	i := 0
	j := n - 1

	for {
		if i > j {
			break
		}

		m := i + (j - i)/2
		if target == nums[m] {
			return m
		}
		if target < nums[m] {
			j = m - 1
		} else if target >nums[m] {
			i = m + 1
		}
	}

	return -1
}

func SearchRange(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1

	n := len(nums)
	if n == 0 {
		return res
	}

	left := 0
	right := n

	for {
		if left >= right {
			break
		}

		mi := (left + right) / 2

		if nums[mi] < target {
			left = mi + 1
		} else {
			right = mi
		}
	}

	if left == n {
		return res
	}

	if nums[left] == target {
		res[0] = left
	} else {
		return res
	}

	right = n
	for {
		if left >= right {
			break
		}

		mi := (left + right) / 2

		if nums[mi] > target {
			right = mi
		} else {
			left = mi + 1
		}
	}

	res[1] = left - 1

	return res
}