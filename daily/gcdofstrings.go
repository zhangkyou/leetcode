package daily

import (
	"strings"
)

func GcdOfStrings(str1 string, str2 string) string {
	n1 := len(str1)
	n2 := len(str2)

	i := 0
	j := 0
	prefix := ""
	for {
		if i >= n1 {
			break
		}

		if j >= n2 {
			break
		}

		if str1[i] == str2[j] {
			prefix += string(str1[i])
			i+=1
			j+=1
		} else {
			break
		}
	}

	for {
		if len(prefix) == 0 {
			break
		}

		if isValid(str1, prefix) && isValid(str2, prefix) {
			return prefix
		} else {
			prefix = prefix[:len(prefix) - 1]
		}
	}

	return ""
}

func isValid(str, prefix string) bool {
	str = strings.Replace(str, prefix, "", -1)
	if len(str) == 0 {
		return true
	}

	return false
}

func GcdOfStringsOptimize(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}

	n := gcd(len(str1), len(str2))

	return str1[:n]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}
