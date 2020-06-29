package graph

import (
	"graph_theory/bag"
)

// Data structure for representing a graph.
type Graph struct {
	nodes int
	infos map[int]interface{}
	edges []*bag.Bag
}

// Creates a new undirected graph.
func New(vertices int) *Graph {
	g := &Graph{
		nodes: vertices,
		infos: make(map[int]interface{}),
		edges: make([]*bag.Bag, vertices),
	}
	for i := 0; i < vertices; i++ {
		g.edges[i] = bag.New()
	}
	return g
}

// Returns the number of vertices in the graph.
func (g *Graph) Vertices() int {
	return g.nodes
}

// Assigns some data to a graph node.
func (g *Graph) Assign(id int, data interface{}) {
	g.infos[id] = data
}

// Retrieves the data associated with a graph node.
func (g *Graph) Retrieve(id int) interface{} {
	return g.infos[id]
}

// Connects two vertices of a graph (may be a loopback).
func (g *Graph) Connect(a, b int) {
	g.edges[a].Insert(b)
	if a != b {
		g.edges[b].Insert(a)
	}
}

// Disconnects two vertices of a graph (may be a loopback).
func (g *Graph) Disconnect(a, b int) {
	g.edges[a].Remove(b)
	if a != b {
		g.edges[b].Remove(a)
	}
}

// Executes a function for every neighbor of a vertex.
func (g *Graph) Do(v int, f func(interface{})) {
	g.edges[v].Do(f)
}
