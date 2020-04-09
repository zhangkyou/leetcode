package weekly

func FindLucky(arr []int) int {
	dict := [500]int{}

	for _, v := range arr {
		dict[v-1] += 1
	}

	for i := 499; i>=0; i-=1 {
		if i == dict[i]-1 {
			return dict[i]
		}
	}
	return -1
}