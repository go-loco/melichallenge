package melichallenge

import (
	"container/list"
	"strconv"
)

type MeliOffice interface {
	ShortestPath() string
}

type office struct {
	me           int
	coffee       int
	candy        int
	chics        map[int]int
	boss         map[int]int
	officeWidth  int
	officeHeight int
	graph        *graph
}

func New(me, coffee, candy Point, chics, boss []Point, officeWidth int, officeHeight int) MeliOffice {

	office := new(office)

	office.officeWidth = officeWidth
	office.officeHeight = officeHeight
	office.me = office.flattenPoint(me)
	office.coffee = office.flattenPoint(coffee)
	office.candy = office.flattenPoint(candy)
	office.chics = office.pToMap(chics)
	office.boss = office.pToMap(boss)

	vertices := officeWidth * officeHeight
	office.graph = newGraph(vertices)
	office.addEdges()

	return office
}

func (o *office) pToMap(s []Point) map[int]int {

	m := make(map[int]int)
	for _, v := range s {
		fp := o.flattenPoint(v)
		m[fp] = fp
	}

	return m
}

func (o *office) iToMap(v int) map[int]int {
	m := make(map[int]int)
	m[v] = v
	return m
}

func (o *office) flattenPoint(p Point) int {
	return ((p.Y * o.officeWidth) + p.X)
}

func (o *office) toPoint(fp int) Point {
	y := int(fp / o.officeWidth)
	x := int(fp - o.officeWidth*y)

	return Point{x, y}
}

func (o *office) addEdges() {

	for p := 0; p < o.graph.vertices; p++ {

		if _, boss := o.boss[p]; !boss {

			point := o.toPoint(p)
			p2 := p + 1

			if _, boss := o.boss[p2]; !boss && point.Y == o.toPoint(p2).Y {
				o.graph.addEdge(p, p2) //; print(p); print("*"); println(p2)
			}

			if point.Y > 0 {
				point3 := Point{point.X, point.Y - 1}
				p3 := o.flattenPoint(point3)

				if _, boss := o.boss[p3]; !boss {
					o.graph.addEdge(p, p3) //print("inside y: "); print(p); print("*"); println(p3)
				}
			}

		} //boss

	}

}

func (o *office) ShortestPath() string {

	targets := new(list.List)
	targets.PushFront(o.chics)
	targets.PushFront(o.iToMap(o.candy))
	targets.PushFront(o.iToMap(o.coffee))

	steps := 1
	start := o.me
	pStart := o.toPoint(start)
	result := pStart.String()

	for t := targets.Front(); t != nil; t = t.Next() {
		target := t.Value.(map[int]int)
		bfs := bfs{o.graph, start, target}
		path := bfs.shortestPath()
		steps += path.Len()

		for e := path.Front(); e != nil; e = e.Next() {
			start = e.Value.(int)
			p := o.toPoint(start)
			result += p.String()
		}

	}

	return "[" + result + "], " + strconv.Itoa(steps)

}
