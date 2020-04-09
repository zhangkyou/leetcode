package util

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a % b)
}