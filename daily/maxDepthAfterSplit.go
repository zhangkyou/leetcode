package daily

func MaxDepthAfterSplit(seq string) []int {
	stack := make([]byte, 0)
	res := make([]int, len(seq))

	for i, c := range seq {
		if c == '(' {
			stack = append(stack, '(')
			res[i] = len(stack)%2
		} else {
			res[i] = len(stack)%2
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

func MaxDepthAfterSplitOptimize(seq string) []int {
	n := 0
	res := make([]int, len(seq))

	for i, c := range seq {
		if c == '(' {
			n += 1
			res[i] = n%2
		} else {
			res[i] = n%2
			n -= 1
		}
	}
	return res
}


