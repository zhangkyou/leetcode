package april

import (
	"fmt"
	"strings"
)

func MyAtoi(str string) int {
	str = strings.Trim(str, " ")
	sLen := len(str)

	if sLen < 1 {
		return 0
	}

	s0 := str
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
		if len(str) < 1 {
			return 0
		}
	}

	max := 0
	if s0[0] == '-' {
		max = 1 << 31
	} else {
		max = 1 << 31 - 1
	}

	n := 0
	for _, ch := range []byte(str) {
		ch -= '0'
		if ch > 9 {
			break
		}
		n = n*10 + int(ch)
		if n >= max {
			n = max
			break
		}
	}
	fmt.Println(n)

	if s0[0] == '-' {
		n = -n
	}

	return n
}
