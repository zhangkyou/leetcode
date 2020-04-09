package april

func Rotate(matrix [][]int)  {
	n := len(matrix)
	if n < 1 {
		return
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			res[j][n-1-i] = matrix[i][j]
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			matrix[i][j] = res[i][j]
		}
	}
}

func RotateOptimize(matrix [][]int)  {
	n := len(matrix)
	if n < 1 {
		return
	}
	for i:=0; i<n/2; i++ {
		for j:=0; j<n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
