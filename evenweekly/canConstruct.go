package evenweekly

func CanConstruct(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}

	mp := [26]uint32{}
	for _, c := range s {
		mp[c-'a'] += 1
	}

	oddNumber := 0
	for _, v := range mp {
		if v%2 == 1 {
			oddNumber += 1
		}
	}

	if oddNumber > k {
		return false
	}
	return true
}
