package array

func RemoveElement(nums []int, val int) int {
	n := len(nums)

	current := 0
	for i:=0; i < n; i+=1 {
		if nums[i] != val {
			nums[current] = nums[i]
			current += 1
		}
	}

	return current
}