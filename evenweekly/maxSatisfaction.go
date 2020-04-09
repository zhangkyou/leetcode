package evenweekly

import (
	"fmt"
	"sort"
)

func MaxSatisfaction(satisfaction []int) int {
	n := len(satisfaction)
	sort.Ints(satisfaction)
	fmt.Println(satisfaction)

	zeroPos := -1
	for i, v := range satisfaction {
		if v >= 0 {
			zeroPos = i
			break
		}
	}
	if zeroPos < 0 {
		return 0
	}
	fmt.Println(zeroPos)

	res := 0
	positiveOnce := 0
	//positive res
	for i:=zeroPos; i<n; i+=1 {
		res += (i-zeroPos+1) * satisfaction[i]
		positiveOnce += satisfaction[i]
	}
	fmt.Println(res)
	fmt.Println(positiveOnce)

	if zeroPos == 0 {
		return res
	}

	max := res
	negativeOnce := 0
	count := 1
	for i:=zeroPos-1; i>=0; i-=1 {
		//negativeOnce += satisfaction[i]
		res += satisfaction[i] + negativeOnce + positiveOnce
		fmt.Println(i, satisfaction[i], res, max, count,negativeOnce, positiveOnce)
		if res > max {
			max = res
			negativeOnce += satisfaction[i]
			count += 1
		}
		fmt.Println(i, satisfaction[i], res, max, count,negativeOnce, positiveOnce)
	}
	return max
}
