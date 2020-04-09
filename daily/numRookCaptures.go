package daily

func NumRookCaptures(board [][]byte) int {
	n := len(board)

	for i:=0; i<n; i+=1 {
		for j:=0; j<n; j+=1 {
			if board[i][j] != 'R' {
				continue
			}
			return capture(board, i, j)
		}
	}
	return 0
}

func capture(board [][]byte, x, y int) int {
	n := len(board)
	res := 0
	for j:=y-1; j>=0; j-=1 {
		if board[x][j] == 'p' {
			res += 1
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	for j:=y+1; j<n; j+=1 {
		if board[x][j] == 'p' {
			res += 1
			break
		} else if board[x][j] == 'B' {
			break
		}
	}

	for i:=x-1; i>=0; i-=1 {
		if board[i][y] == 'p' {
			res += 1
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	for i:=x+1; i<n; i+=1 {
		if board[i][y] == 'p' {
			res += 1
			break
		} else if board[i][y] == 'B' {
			break
		}
	}

	return res
}
