package meli

import "container/list"

type bfs struct {
	graph          *graph
	initVertice    int
	targetVertices map[int]int
}

func (b *bfs) shortestPath() *list.List {

	marked := make([]bool, b.graph.vertices)
	edgeTo := make([]int, b.graph.vertices)
	queue := new(queue)
	targetFound := false
	var target int

	queue.enqueue(b.initVertice)
	marked[b.initVertice] = true

	for queue.size > 0 && !targetFound {
		currentVertice := queue.dequeue().(int)
		adjacents := b.graph.adjacents(currentVertice)

		for e := adjacents.Front(); e != nil; e = e.Next() {
			v := e.Value.(int)
			if !marked[v] {
				marked[v] = true
				edgeTo[v] = currentVertice
				queue.enqueue(v)
			}

			if _, ok := b.targetVertices[v]; ok {
				targetFound = true
				target = v
			}

		}

	}

	result := new(list.List)
	if targetFound {
		result.PushFront(target)
		for v := edgeTo[target]; v != b.initVertice; v = edgeTo[v] {
			result.PushFront(v)
		}
	}

	return result

}
