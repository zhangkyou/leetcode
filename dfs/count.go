package dfs

func numIslands(grid [][]byte) int {
	counter := 0
	for i:=0; i<len(grid); i+=1 {
		for j:=0; j<len(grid[0]); j+=1 {
			if string(grid[i][j]) == "1" {
				area := getArea2(grid, i, j)
				if area > 0 {
					counter += 1
				}
			}
		}
	}
	return counter
}

func getArea2(grid [][]byte, i, j int) int {
	if i >= len(grid) || i < 0 {
		return 0
	}

	if j >= len(grid[0]) || j < 0 {
		return 0
	}

	if string(grid[i][j]) == "1" {
		grid[i][j] = 0
		return 1 + getArea2(grid, i+1, j) + getArea2(grid, i, j+1) + getArea2(grid, i-1, j) + getArea2(grid, i, j-1)
	}

	return 0
}
