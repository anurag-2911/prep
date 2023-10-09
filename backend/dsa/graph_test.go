package dsa

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	fmt.Println("test")
}

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}
