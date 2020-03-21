package reverse

import (
	"fmt"
	"math"
)

func IntReverse(x int) int {
	var temp int
	temp = 0

	for ; x != 0; x /= 10{
		pop := x%10
		temp = temp * 10 + pop
		if temp > math.MaxInt32 || temp < math.MinInt32 {
			return 0
		}
	}

	return temp
}

func IsEqual(x int) bool {
	temp := IntReverse(x)
	fmt.Println(x)
	fmt.Println(temp)
	return temp == x
}

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x !=0) {
		return false
	}
	reversenum := 0
	for{
		if x <= reversenum {
			break
		}
		reversenum = reversenum * 10 + x % 10;
		x /= 10
	}
	fmt.Println(reversenum)
	fmt.Println(x)

	return x == reversenum || x == reversenum/10
}
