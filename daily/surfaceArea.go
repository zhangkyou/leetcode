package daily

import (
	"Goproject/leetcode/util"
)

func SurfaceArea(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	res := 0
	for i:=0; i<n; i+=1 {
		for j:=0; j<n; j+=1 {
			if grid[i][j] == 0 {
				continue
			}
			cur := 10 + 4*(grid[i][j] - 2)
			//left
			if i > 0 {
				cur -= util.Min(grid[i-1][j], grid[i][j]) << 1
			}
			//right
			//if i < n-1 {
			//	cur -= min(grid[i+1][j], grid[i][j])
			//}
			//up
			if j > 0 {
				cur -= util.Min(grid[i][j-1], grid[i][j]) << 1
			}
			//down
			//if j < n-1 {
			//	cur -= min(grid[i][j+1], grid[i][j])
			//}
			res += cur
		}
	}
	return res
}
