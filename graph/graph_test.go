package graph

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {
	matrix := [][]int{
		{3, 2, 1},
		{2, 4, 3},
		{2, 1, 3},
		{3, 2, 4},
	}

	g := CreateGraph(matrix)

	BFS(g.Nodes[2])
}

func TestSortTopology(t *testing.T) {
	matrix := [][]int{
		{1, 1, 2},
		{1, 1, 3},
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
	}

	g := CreateGraph(matrix)

	SortTopology(g)
}