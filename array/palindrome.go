package array

func Palindrome(input string) bool {
	mp := make(map[string]int)

	for _, s := range input {
		if _, ok := mp[string(s)]; ok {
			mp[string(s)] += 1
		} else {
			mp[string(s)] = 1
		}
	}

	count := 0
	for _, v := range mp {
		if v % 2 == 1 {
			count += 1
		}
	}

	if count == 0 || count == 1 {
		return true
	}
	return false
}
