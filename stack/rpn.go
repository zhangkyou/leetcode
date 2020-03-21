package stack

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, s := range tokens {
		switch s {
		case "+":
			first := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			tmp := first + second
			stack = append(stack, tmp)
		case "-":
			first := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			tmp := second - first
			stack = append(stack, tmp)
		case "/":
			first := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			tmp := second/first
			stack = append(stack, tmp)
		case "*":
			first := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			second := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			tmp := first*second
			stack = append(stack, tmp)
		default:
			number, _ := strconv.Atoi(s)
			stack = append(stack, number)
		}
	}

	return stack[0]
}