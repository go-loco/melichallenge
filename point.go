package melichallenge

import "strconv"

type Point struct {
	X, Y int
}

func (p *Point) Equal(point Point) bool {
	if point.X == p.X && point.Y == p.Y {
		return true
	}

	return false
}

func (p *Point) String() string {
	return ("(" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + ")")
}
