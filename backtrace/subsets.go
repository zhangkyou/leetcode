package backtrace

var res [][]int
var temp []int

var path []int
var used []bool

func Backtrace(nums []int) [][]int {
	res = make([][]int, 0)
	temp = make([]int, 0)

	trace(nums, 0)
	return res
}

func trace(nums []int, begin int) {
	add := make([]int, len(temp))
	copy(add, temp)
	res = append(res, add)

	for i:=begin;i < len(nums); i+=1 {
		temp = append(temp, nums[i])

		trace(nums, i+1)
		temp = temp[:len(temp) - 1]
	}
}
