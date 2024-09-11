package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/diegoholiveira/go-graph"
)

func TestShortestPath(t *testing.T) {
	graph := graph.New()
	graph.AddEdge("A", "B", 2)
	graph.AddEdge("A", "C", 6)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 8)
	graph.AddEdge("D", "E", 10)
	graph.AddEdge("D", "F", 15)
	graph.AddEdge("E", "F", 6)
	graph.AddEdge("E", "G", 2)
	graph.AddEdge("F", "G", 6)

	path := graph.ShortestPath("A", "G")

	expectedNodes := []string{"A", "B", "D", "E", "G"}
	expectedDistance := 19

	assert.Equal(t, expectedNodes, path.Nodes())
	assert.Equal(t, expectedDistance, path.Distance())
}
