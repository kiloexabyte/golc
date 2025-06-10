package mathandgeometry

// Leetcode #2013
type DetectSquares struct {
	points map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points: make(map[[2]int]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	key := [2]int{point[0], point[1]}
	this.points[key]++
}

func (this *DetectSquares) Count(point []int) int {
	px, py := point[0], point[1]
	res := 0

	for p, count := range this.points {
		x, y := p[0], p[1]
		if abs(px-x) == abs(py-y) && (px != x || py != y) {
			m1 := this.points[[2]int{x, py}]
			m2 := this.points[[2]int{px, y}]
			res += m1 * m2 * count
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */