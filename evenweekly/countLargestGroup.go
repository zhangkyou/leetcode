package evenweekly

func CountLargestGroup(n int) int {
	mp := make(map[int]int)

	max := 0
	for i:=1; i<=n; i+=1 {
		key := sumBit(i)
		mp[key] += 1
		if mp[key] > max {
			max = mp[key]
		}
	}

	count := 0
	for _, v := range mp {
		if max == v {
			count += 1
		}
	}
	return count
}

func sumBit(n int) int {
	res := 0
	for {
		if n == 0 {
			break
		}
		res += n%10
		n /= 10
	}
	return res
}
