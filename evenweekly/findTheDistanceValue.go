package evenweekly

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := 0; i<len(arr1); i+=1 {
		flag := false
		for j:=0; j<len(arr2); j+=1 {
			if abs(arr1[i] - arr2[j]) <= d {
				flag = true
				break
			}
		}
		if flag == false {
			res += 1
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
