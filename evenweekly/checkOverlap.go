package evenweekly

import (
	"math"
)

func CheckOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	//fmt.Println(x_center - radius, x2)
	//fmt.Println(x_center + radius, x1)
	//fmt.Println(y_center + radius, y1)
	//fmt.Println(y_center - radius, y2)
	if (x_center >= x1 && x_center <= x2) || (y_center >= y1 && y_center <= y2) {
		if x_center-radius > x2 || x_center+radius < x1 || y_center+radius < y1 || y_center-radius > y2 {
			return false
		}
	} else if x_center > x2 && y_center < y1{ //right down
		//fmt.Println(x_center, x2)
		//fmt.Println(y_center, y1)
		dis := distance(x_center, y_center, x2, y1)
		//fmt.Println(dis)
		if dis <= float64(radius) {
			return true
		} else {
			return false
		}
	} else if x_center < x1 && y_center < y1 { //left down
		dis := distance(x_center, y_center, x1, y1)
		if dis <= float64(radius) {
			return true
		} else {
			return false
		}
	} else if x_center < x1 && y_center > y2 { //left up
		dis := distance(x_center, y_center, x1, y2)
		if dis <= float64(radius) {
			return true
		} else {
			return false
		}
	} else if x_center > x2 && y_center > y2 { //right up
		dis := distance(x_center, y_center, x2, y2)
		if dis <= float64(radius) {
			return true
		} else {
			return false
		}
	}

	return true
}

func distance(x_center, y_center, x, y int) float64 {
	x_distance := x - x_center
	y_distance := y - y_center
	return math.Sqrt(float64(x_distance * x_distance + y_distance * y_distance))
}
