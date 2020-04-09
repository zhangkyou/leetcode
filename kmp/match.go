package kmp

func ViolentMatch(s, p string) int {
	sn := len(s)
	pn := len(p)

	i := 0
	j := 0
	for {
		if i >= sn || j >= pn {
			break
		}

		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == pn {
		return i - j
	}
	return -1
}

func KmpSearch(s, p string) int {
	i := 0
	j := 0
	sn := len(s)
	pn := len(p)

	next := getNext(p)
	for {
		if i >= sn || j >= pn {
			break
		}

		//①如果j = -1，或者当前字符匹配成功（即S[i] == P[j]），都令i++，j++
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			//②如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令 i 不变，j = next[j]
			//next[j]即为j所对应的next值
			j = next[j]
		}
	}
	if j == pn {
		return i - j
	}
	return -1
}

func getNext(input string) []int {
	n := len(input)
	next := make([]int, n)
	next[0] = -1
	k := -1
	j := 0
	for {
		if j >= n-1 {
			break
		}
		if k == -1 || input[j] == input[k] {
			j++
			k++
			if input[j] != input[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
	return next
}
