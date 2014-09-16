package meli

import ("testing")

func TestMeli1(t *testing.T) {

  me := Point{1,0}
  coffee := Point{6,1}
  candy := Point{1,6}
  chics := []Point{ Point{5,5}, Point{5,6}, Point{6,5}, Point{6,6} }
  boss := []Point{ Point{1,2}, Point{2,2}, Point{2,3}, Point{2,4}, Point{3,2}, Point{3,3}, Point{3,4}, Point{4,2} }
  N, M := 8, 8

  meliOffice := NewMeliOffice(me, coffee, candy, chics, boss, N, M)

  println(meliOffice.ShortestPath())
}

func TestMeli2(t *testing.T) {

  me := Point{0,0}
  coffee := Point{6,1}
  candy := Point{1,7}
  chics := []Point{ Point{5,5}, Point{5,6}, Point{6,5}, Point{6,6} }
  boss := []Point{ Point{1,2} }
  N, M := 8, 8

  meliOffice := NewMeliOffice(me, coffee, candy, chics, boss, N, M)

  println(meliOffice.ShortestPath())
}

func TestMeli3(t *testing.T) {

  me := Point{1,0}
  coffee := Point{6,1}
  candy := Point{6,2}
  chics := []Point{ Point{5,5}, Point{5,6}, Point{6,5}, Point{6,6} }
  boss := []Point{ Point{1,2} }//, Point{2,2}, Point{2,3}, Point{2,4}, Point{3,2}, Point{3,3}, Point{3,4}, Point{4,2} }
  N, M := 8, 8

  meliOffice := NewMeliOffice(me, coffee, candy, chics, boss, N, M)

  println(meliOffice.ShortestPath())
}