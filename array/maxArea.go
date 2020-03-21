package array

func MaxArea(height []int) int {
	area := 0
	n := len(height)
	left := 0
	right := n - 1
	for {
		if left >= right {
			break
		}
		currentArea := min(height[left], height[right]) * (right - left)
		if currentArea > area {
			area = currentArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
