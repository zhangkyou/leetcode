package stack

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
 	var stack = make([]string, 0)

	for _, c := range s {
		switch string(c) {
		case ")":
			if len(stack) == 0 || stack[len(stack) - 1] != "(" {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		case "]":
			if len(stack) == 0 || stack[len(stack) - 1] != "[" {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		case "}":
			if len(stack) == 0 || stack[len(stack) - 1] != "{" {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		default:
			stack = append(stack, string(c))
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
