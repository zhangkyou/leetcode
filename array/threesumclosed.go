package array

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	mindiff := math.MaxInt8
	res := 0
	if n <= 3 {
		for _, num := range nums {
			res += num
		}
		return res
	}

	sort.Ints(nums)

	for i:=0; i < n - 2; i++ {
		left := i + 1
		right := n - 1
		if i>0 && nums[i] == nums[i - 1] {
			continue
		}

		for {
			if left >= right {
				break
			}

			current := nums[i] + nums[left] + nums[right]
			if current - target == 0 {
				return current
			}

			if abs(current - target) < abs(mindiff) {
				mindiff = current - target
				res = current
			}

			if current - target < 0 {
				left++
				for ; nums[left] == nums[left-1]; left++ {
					if left >= right {
						break
					}
				}
			}

			if current - target > 0 {
				right--
				for ; nums[right] == nums[right+1]; right-- {
					if left >= right {
						break
					}
				}
			}
		}
	}

	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}