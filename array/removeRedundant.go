package array

func RemoveDuplicates(nums []int) int {
	n := len(nums)

	last := 1

	for i := 1; i < n ; i+=1 {
		if nums[i] != nums[i - 1] {
			nums[last] = nums[i]
			last++
		}
	}
	return last
}