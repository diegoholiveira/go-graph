package graph

import (
	"container/heap"
	"math"
)

type Graph struct {
	edges    map[string]map[string]int
	vertexes []string
}

type Path struct {
	nodes    []string
	distance int
}

func New() *Graph {
	return &Graph{
		edges:    make(map[string]map[string]int),
		vertexes: make([]string, 0),
	}
}

func (g *Graph) AddEdge(src, dest string, distance int) {
	if _, found := g.edges[src]; !found {
		g.edges[src] = make(map[string]int)
		g.vertexes = append(g.vertexes, src)
	}

	if _, found := g.edges[dest]; !found {
		g.edges[dest] = make(map[string]int)
		g.vertexes = append(g.vertexes, dest)
	}

	g.edges[src][dest] = distance
}

func (g *Graph) ShortestPath(src, dest string) *Path {
	distances := make(map[string]int)
	for _, vertex := range g.vertexes {
		distances[vertex] = math.MaxInt
	}

	distances[src] = 0

	toVisit := make(vertexesQueue, 0)
	previous := make(map[string]string)

	heap.Init(&toVisit)
	heap.Push(&toVisit, nextVertex{distance: 0, current: src})

	for toVisit.Len() > 0 {
		next := heap.Pop(&toVisit).(nextVertex)
		if next.distance > distances[next.current] {
			continue
		}

		for vertex, distance := range g.edges[next.current] {
			total := next.distance + distance

			if distances[vertex] >= total {
				distances[vertex] = total

				previous[vertex] = next.current

				heap.Push(&toVisit, nextVertex{
					distance: total,
					current:  vertex,
				})
			}
		}
	}

	path := &Path{
		nodes:    reconstructPath(previous, dest),
		distance: distances[dest],
	}

	return path
}

func (p *Path) Nodes() []string {
	return p.nodes
}

func (p *Path) Distance() int {
	return p.distance
}

func reconstructPath(previous map[string]string, target string) []string {
	path := []string{}
	for target != "" {
		path = append([]string{target}, path...)
		target = previous[target]
	}
	return path
}
