package meli

import("strconv")

type Point struct {
  x, y int
}

func (p *Point) Equal(point Point) bool { 
  if point.x == p.x && point.y == p.y {
    return true 
  } else {
    return false
  }
}

func (p *Point) toString() string {
  return ("(" + strconv.Itoa(p.x) + ", " + strconv.Itoa(p.y) + ")")
}