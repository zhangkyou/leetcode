package redundent

//维护滑动窗口的做边界值
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	subMap := make(map[byte]int)

	for i,j := 0, 0; j < n; j++ {
		if v, ok := subMap[byte(s[j])]; ok {
			if i < v {
				i = v
			}
 		}
		if ans < (j - i + 1) {
			ans = j - i + 1
		}
		subMap[byte(s[j])] = j + 1
	}


	return ans
}
