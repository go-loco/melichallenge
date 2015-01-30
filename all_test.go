package melichallenge

import "testing"

func TestMeli1(t *testing.T) {

	me := Point{1, 0}
	coffee := Point{6, 1}
	candy := Point{1, 6}
	chics := []Point{Point{5, 5}, Point{5, 6}, Point{6, 5}, Point{6, 6}}
	boss := []Point{Point{1, 2}, Point{2, 2}, Point{2, 3}, Point{2, 4}, Point{3, 2}, Point{3, 3}, Point{3, 4}, Point{4, 2}}
	N, M := 8, 8

	meliOffice := New(me, coffee, candy, chics, boss, N, M)

	println(meliOffice.ShortestPath())
}

func TestMeli2(t *testing.T) {

	me := Point{0, 0}
	coffee := Point{6, 1}
	candy := Point{1, 7}
	chics := []Point{Point{5, 5}, Point{5, 6}, Point{6, 5}, Point{6, 6}}
	boss := []Point{Point{1, 2}}
	N, M := 8, 8

	meliOffice := New(me, coffee, candy, chics, boss, N, M)

	println(meliOffice.ShortestPath())
}

func TestMeli3(t *testing.T) {

	me := Point{1, 0}
	coffee := Point{6, 1}
	candy := Point{6, 2}
	chics := []Point{Point{5, 5}, Point{5, 6}, Point{6, 5}, Point{6, 6}}
	boss := []Point{Point{1, 2}} //, Point{2,2}, Point{2,3}, Point{2,4}, Point{3,2}, Point{3,3}, Point{3,4}, Point{4,2} }
	N, M := 8, 8

	meliOffice := New(me, coffee, candy, chics, boss, N, M)

	println(meliOffice.ShortestPath())
}

func TestMeli4(t *testing.T) {

	me := Point{0, 0}
	coffee := Point{9, 1}
	candy := Point{6, 4}
	chics := []Point{Point{14, 8}, Point{14, 9}, Point{15, 8}, Point{15, 9}}
	boss := []Point{
		Point{1, 2}, Point{2, 2}, Point{3, 2}, Point{4, 2}, Point{5, 2}, Point{6, 2}, Point{7, 2}, Point{8, 2}, Point{9, 2}, Point{10, 2}, Point{11, 2}, Point{12, 2},
		Point{1, 3}, Point{2, 3}, Point{3, 3}, Point{4, 3}, Point{5, 3}, Point{6, 3}, Point{7, 3}, Point{8, 3}, Point{9, 3}, Point{10, 3}, Point{11, 3}, Point{12, 3},
		Point{1, 4}, Point{2, 4}, Point{3, 4}, Point{4, 4}, Point{5, 4}, Point{7, 4}, Point{8, 4}, Point{9, 4}, Point{10, 4}, Point{11, 4}, Point{12, 4},
		Point{1, 5}, Point{2, 5}, Point{3, 5}, Point{4, 5}, Point{5, 5}, Point{7, 5}, Point{8, 5}, Point{9, 5}, Point{10, 5}, Point{11, 5}, Point{12, 5},
	}
	N, M := 16, 15

	meliOffice := New(me, coffee, candy, chics, boss, N, M)

	println(meliOffice.ShortestPath())
}
