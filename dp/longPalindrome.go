package dp

import (
	"fmt"
)

func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	if n == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	last := string(s[0])
	var temp string
	for i:= 1; i<n-1; i++ {
		if s[i-1] == s[i+1] {
			temp = search(s, i)
		}
		if len(temp) > len(last) {
			last = temp
		}
	}

	for i:=0; i<n-1; i++ {
		if s[i] == s[i+1] {
			temp = search2(s, i, i+1)
		}
		if len(temp) > len(last) {
			last = temp
		}
	}
	return last
}

func search(s string, i int) string {
	n := len(s)
	offset := 1
	for {
		if i - offset < 0 || i + offset > n - 1 {
			offset--
			break
		}
		if s[i - offset] != s[i + offset] {
			offset--
			break
		}
		offset++
	}

	return s[i - offset : i + offset + 1]
}

func search2(s string, left int, right int) string {
	n := len(s)
	offset := 0
	for {
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(offset)
		if left - offset < 0 || right + offset > n - 1 {
			offset--
			break
		}
		if s[left - offset] != s[right + offset] {
			offset--
			break
		}
		offset++
	}
	return s[left - offset : right+offset+1]
}
