package array

func IsValidSudoku(board [][]byte) bool {
	n := len(board)

	rows := make([]map[byte]int, n)
	columns := make([]map[byte]int, n)
	block := make([]map[byte]int, n)

	for i:=0; i<n; i+=1 {
		rows[i] = make(map[byte]int)
		columns[i] = make(map[byte]int)
		block[i] = make(map[byte]int)
	}

	for i:=0; i<n; i+=1 {
		for j:=0; j<n; j+=1 {
			if string(board[i][j]) == "." {
				continue
			}
			if _, ok := rows[i][board[i][j]]; ok {
				return false
			} else {
				rows[i][board[i][j]] = 1
			}

			if _, ok := columns[j][board[i][j]]; ok {
				return false
			} else {
				columns[j][board[i][j]] = 1
			}

			blockid := i/3 * 3 + j/3
			if _, ok := block[blockid][board[i][j]]; ok {
				return false
			} else {
				block[blockid][board[i][j]] = 1
			}
		}
	}

	return true
}