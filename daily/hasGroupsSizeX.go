package daily

import (
	"Goproject/leetcode/util"
	"sort"
)

func HasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	dict := make([]int, 10000)

	for i := range deck {
		dict[deck[i]] += 1
	}

	sort.Ints(dict)
	one := 0
	two := 0
	for i:=0; i<len(dict); i+=1 {
		if one == 0 && dict[i] > 0 {
			one = dict[i]
		}
		if one > 0 && dict[i] > 0 && dict[i] != one {
			two = dict[i]
			break
		}
	}

	gcd := util.Gcd(one, two)
	//fmt.Println(dict, one, two, gcd)
	if gcd < 2 {
		return false
	}

	for i:=0; i<len(dict); i+=1 {
		if dict[i] % gcd != 0 {
			return false
		}
	}

	return true
}
