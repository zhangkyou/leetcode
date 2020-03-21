package roman

func RomanToInt(s string) int {
	dict := map[string]int{
		"I" : 1,
		"V" : 5,
		"X"	: 10,
		"L"	: 50,
		"C"	: 100,
		"D"	: 500,
		"M"	: 1000,
	}

	n := len(s)
	result := 0
	for i:=0; i<n-1; i++ {
		if value, ok := dict[string(s[i])]; ok {
			if string(s[i]) == "I" {
				if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
					result -= dict["I"]
				} else {
					result += dict["I"]
				}
			} else if string(s[i]) == "X" {
				if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
					result -= dict["X"]
				} else {
					result += dict["X"]
				}
			} else if string(s[i]) == "C" {
				if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
					result -= dict["C"]
				} else {
					result += dict["C"]
				}
			} else {
				result += value
			}
		}
	}

	result += dict[string(s[n-1])]
	return result
}
