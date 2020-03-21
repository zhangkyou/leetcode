package array

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	for i:=0; i < n - 2; i++ {
		if nums[i] > 0 {
			break
		}
		left := i + 1
		right := n - 1
		if nums[i] * nums[right] > 0 {
			break
		}
		if i>0 && nums[i] == nums[i - 1] {
			continue
		}

		for {
			if left >= right {
				break
			}

			current := nums[i] + nums[left] + nums[right]
			if current == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				for ; nums[left] == nums[left-1]; left++ {
					if left >= right {
						break
					}
				}
				right--
				for ; nums[right] == nums[right+1]; right-- {
					if left >= right {
						break
					}
				}
			} else if current < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
