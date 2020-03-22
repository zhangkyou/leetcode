package daily

import (
	"strconv"
)

func CompressString(S string) string {
	n := len(S)
	if n == 0 {
		return S
	}
	dict := [52]int{}

	var last uint8
	if S[0] < 'a' {
		last = S[0] - 'A'
	} else {
		last = S[0] - 'a' + 'Z' - 'A' + 1
	}
	var k uint8
	var res string
	for i := range S {
		if S[i] < 'a' {
			k = S[i] - 'A'
		} else {
			k = S[i] - 'a' + 'Z' - 'A' + 1
		}
		if last == k {
			dict[k] += 1
		} else {
			if last < 26 {
				res += string(last + 'A')
			} else {
				res += string(last - ('Z' - 'A' + 1) + 'a')
			}
			res += strconv.Itoa(dict[last])
			dict[last] = 0
			last = k
			dict[k] = 1
		}
	}
	if last < 26 {
		res += string(last + 'A')
	} else {
		res += string(last - ('Z' - 'A' + 1) + 'a')
	}
	res += strconv.Itoa(dict[last])

	if len(res) < n {
		return res
	}
	return S
}
