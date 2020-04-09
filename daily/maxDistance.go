package daily

type node struct{
	x,y,s int
}

func MaxDistance(grid [][]int) int {
	var q [10000]node
	t := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	h := 0
	l := h
	ll := len(grid[0])
	max := -1
	for i, v := range grid {
		for j, w := range v {
			if w == 1 {
				w = 0
				q[l].x = i
				q[l].y = j
				q[l].s = 0
				l++
			}
		}
	}
	for h < l {
		for i := 0; i < 4; i++ {
			x := q[h].x + t[i][0]
			y := q[h].y + t[i][1]
			s := q[h].s
			if ll <= x || ll <= y || x < 0 || y < 0 || grid[x][y] != 0 {
				continue
			}
			grid[x][y] = s + 1
			q[l].x = x
			q[l].y = y
			q[l].s = s + 1
			l++
			max = s + 1
		}
		h++
	}
	return max
}