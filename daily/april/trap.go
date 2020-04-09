package april

import "Goproject/leetcode/util"

func Trap(height []int) int {
	max := 0
	for _, v := range height {
		if v > max {
			max = v
		}
	}

	res := 0
	for i:=1; i <= max; i+=1 {
		res += get(height, i)
	}
	return res
}

func get(height []int, object int) int {
	start := 0
	for i:=0; i<len(height); i+=1 {
		if height[i] >= object {
			start = i
			break
		}
	}
	end := 0
	for i:=len(height)-1; i>=0; i-=1 {
		if height[i] >= object {
			end = i
			break
		}
	}

	res := 0
	for i:=start; i<end; i+=1 {
		if height[i] < object {
			res += 1
		}
	}
	return res
}

func TrapOptimize(height []int) int {
	leftHigh := make([]int, len(height))
	rightHigh := make([]int, len(height))

	for i:=0; i<len(leftHigh); i+=1 {
		if i == 0 {
			leftHigh[i] = 0
			continue
		}
		leftHigh[i] = util.Max(leftHigh[i-1], height[i-1])
	}

	for i:=len(rightHigh) - 1; i>=0; i-=1 {
		if i == len(rightHigh) - 1 {
			rightHigh[i] = 0
			continue
		}
		rightHigh[i] = util.Max(rightHigh[i+1], height[i+1])
	}

	res := 0
	for i:=0; i<len(height); i+=1 {
		tmp := util.Min(leftHigh[i], rightHigh[i]) - height[i]
		if tmp > 0 {
			res += tmp
		}
	}
	return res
}
