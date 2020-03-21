package commonstring

func LongestCommonPrefix(strs []string) string {
	strslen := len(strs)

	if strslen == 0 {
		return ""
	}

	if strslen < 2 {
		return strs[0]
	}
	substring := getCommonString(strs[0], strs[1])

	if strslen > 2 && len(substring) > 0 {
		for i:=2; i < strslen; i++ {
			substring = getCommonString(substring, strs[i])
		}
	}

	return substring
}

func getCommonString(first string, second string) string {
	firstlen := len(first)
	secondlen := len(second)
	l := firstlen
	if secondlen < firstlen {
		l = secondlen
	}
	samelen := 0
	for i:=0; i < l; i++ {
		if first[i] == second[i] {
			samelen++
		} else {
			break
		}
	}

	return first[:samelen]
}

