package dfs

func MaxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i:=0; i<len(grid); i+=1 {
		for j:=0; j<len(grid[0]); j+=1 {
			if grid[i][j] == 1 {
				area := getArea(grid, i, j)
				maxArea = max(area, maxArea)
			}
		}
	}
	return maxArea
}

func getArea(grid [][]int, i, j int) int {
	if i >= len(grid) || i < 0 {
		return 0
	}

	if j >= len(grid[0]) || j < 0 {
		return 0
	}

	if grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + getArea(grid, i+1, j) + getArea(grid, i, j+1) + getArea(grid, i-1, j) + getArea(grid, i, j-1)
	}

	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}