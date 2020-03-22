package daily

func CanMeasureWater(x int, y int, z int) bool {
	if x + y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x + y == z
	}

	return z % gcd(x, y) == 0
}
