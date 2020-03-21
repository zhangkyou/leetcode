package daily

import (
	"fmt"
)

func LongestPalindrome(s string) int {
	//用数组比map更省空间
	mp := make(map[int32]int)

	for _, c := range s {
		if _, ok := mp[c]; ok {
			mp[c] += 1
		} else {
			mp[c] = 1
		}
	}

	res := 0
	isHasOdd := false
	for _, v := range mp {
		if v % 2 == 0 {
			res += v
		} else {
			isHasOdd = true
			res += v - 1
		}
	}

	if isHasOdd {
		res += 1
	}
	return res
}

func LongestPalindromeOptimize(s string) int {
	dict := [52]int{}

	for i := range s {
		var k uint8
		if s[i] < 'a' {
			k = s[i] - 'A'
		} else {
			k = s[i] - 'a' + 'Z' - 'A' + 1
		}
		dict[k] += 1
	}

	res := 0
	for _, v := range dict {
		fmt.Println(v)
		if v % 2 == 0 {
			res += int(v)
		} else {
			res += int(v - 1)
		}
	}

	if len(s) > res {
		res += 1
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
