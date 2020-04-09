package april

func GameOfLife(board [][]int)  {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	middle := make([][]int, m)
	for i := range middle {
		middle[i] = make([]int, n)
	}
	for i:=0; i<m; i+=1 {
		for j:=0; j<n; j+=1 {
			status := isAlive(board, i, j)
			if status == -1 {
				middle[i][j] = 0
			} else if status == 1 {
				middle[i][j] = 1
			}
		}
	}

	for i:=0; i<m; i+=1 {
		for j:=0; j<n; j+=1 {
			board[i][j] = middle[i][j]
		}
	}
}

func isAlive(board [][]int, x, y int) int {
	check := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1}}

	aliveCount := 0
	for i:=0; i<len(check); i+=1 {
		checkX := x + check[i][0]
		checkY := y + check[i][1]
		if checkX >= 0 && checkY >= 0 && checkX < len(board) && checkY < len(board[0]) {
			if board[checkX][checkY] == 1 {
				aliveCount += 1
			}
		}
	}

	if board[x][y] == 1 {
		if aliveCount < 2 {
			return -1
		}
		if aliveCount == 2 || aliveCount == 3 {
			return 1
		}
		if aliveCount > 3 {
			return -1
		}
	} else {
		if aliveCount == 3 {
			return 1
		}
	}
	return 0
}
