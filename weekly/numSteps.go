package weekly

func NumSteps(s string) int {
	str := []byte(s)
	n := len(str)
	res := 0
	i := n - 1
	for {
		if i <= 0 {
			break
		}

		if str[i] == '0' {
			res++
			i--
		} else {
			res++
			for {
				if i < 0 || str[i] == '0' {
					break
				}
				res++
				i--
			}
			if i>0 {
				str[i] = '1'
			}
		}
	}
	return res
}