package binary

func FindPeakElement(nums []int) int {
	return binarySearch(nums, 0, len(nums) - 1)
}

func binarySearch(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] {
		return binarySearch(nums, left, mid)
	}

	return binarySearch(nums, mid+1, right)
}

func FindPeakElementFor(nums []int) int {
	n := len(nums)

	left := 0
	right := n - 1

	for {
		if left >= right {
			break
		}

		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}