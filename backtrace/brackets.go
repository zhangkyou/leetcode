package backtrace

var bracketsRes []string

func GenerateParenthesis(n int) []string {
	bracketsRes = make([]string, 0)

	if n == 0 {
		return bracketsRes
	}

	dfsBrackets("", n, n)
	return bracketsRes
}


func dfsBrackets(cur string, left, right int) {
	if left == 0 && right == 0 {
		bracketsRes = append(bracketsRes, cur)
		return
	}


	if left > right {
		return
	}

	if left > 0 {
		dfsBrackets(cur + "(", left - 1, right)
	}

	if right > 0 {
		dfsBrackets(cur + ")", left, right - 1)
	}
}