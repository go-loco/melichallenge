package meli

import (
  "container/list"
  "strconv"
)

type MeliOffice interface {
  ShortestPath() string
}

type office struct {
  me     int
  coffee int
  candy  int
  chics  map[int]int
  boss   map[int]int
  officeWidth int
  officeHeight int
  graph *graph
}

func NewMeliOffice(me, coffee, candy Point, chics, boss []Point, officeWidth int, officeHeight int) MeliOffice {

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
  return (p.y * o.officeWidth + p.x)
}

func (o *office) toPoint(fp int) Point {
  y := int( fp / o.officeHeight )
  x := int( fp - o.officeWidth * y )

  return Point{x, y}
}

func (o *office) addEdges() {
  
    for p := 0; p < o.graph.vertices; p++ {

      if _, boss := o.boss[p]; !boss {

        p2 := p + 1
        y := o.toPoint(p).y

        if _, boss := o.boss[p2]; !boss && y == o.toPoint(p2).y {
          o.graph.addEdge(p, p2) //println(p1.toString() + p2.toString())
        }

        if y > 0 {
          p3 := p - o.officeWidth
          if _, boss := o.boss[p3]; !boss {
            o.graph.addEdge(p, p3) //println(p1.toString() + p3.toString())
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
  result := pStart.toString()

  for t := targets.Front(); t != nil; t = t.Next() {
    target := t.Value.(map[int]int)
    bfs := bfs{ o.graph, start, target }
    path := bfs.shortestPath()
    steps += path.Len()

    for e := path.Front(); e != nil; e = e.Next() {
      start = e.Value.(int)
      p := o.toPoint(start)
      result += p.toString()
    }
    

  }

  return "[" + result + "], " + strconv.Itoa(steps)

}

