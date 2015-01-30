package melichallenge

import "container/list"

type graph struct {
	edges    int
	vertices int
	adj      []list.List
}

func newGraph(vertices int) *graph {
	return &graph{
		0,
		vertices,
		make([]list.List, vertices),
	}
}

func (g *graph) addEdge(v, w int) {
	g.adj[v].PushBack(w)
	g.adj[w].PushBack(v)
	g.edges++
}

func (g *graph) adjacents(vertice int) *list.List {
	return &g.adj[vertice]
}
