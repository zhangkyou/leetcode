package daily

var mp []int

/*
hash冲突的线性探测
 */
func MinIncrementForUnique(A []int) int {
	mp = make([]int, 80000)
	for i := range mp {
		mp[i] = -1
	}

	res := 0
	for _, v := range A {
		b := findPos(v)
		res += b - v
	}
	return res
}

func findPos(a int) int {
	b := mp[a]
	if b == -1 {
		mp[a] = a
		return a
	}
	b = findPos(b+1)
	mp[a] = b
	return b
}
