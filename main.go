package main

import (
	"fmt"
	"graph_theory/graph"

	"graph_theory/graph/bfs"
	"graph_theory/graph/dfs"
)

// 6 8
// 1 2
// 2 3
// 2 4
// 3 4
// 3 5
// 3 6
// 4 6
// 5 6
// 1 5

func main() {
	// Create the graph
	g := graph.New(8)
	g.Connect(1, 2)
	g.Connect(2, 3)
	g.Connect(2, 4)
	g.Connect(3, 4)
	g.Connect(3, 5)
	g.Connect(3, 6)
	g.Connect(4, 6)
	g.Connect(5, 6)

	// Create the depth first search algo structure for g and source node #2
	d := dfs.New(g, 1)

	// Get the path between #0 (source) and #2
	fmt.Println("Path 0->5:", d.Path(5))
	fmt.Println("Order:", d.Order())
	fmt.Println("Reachable #4 #6:", d.Reachable(4), d.Reachable(6))
	fmt.Println("Reachable #7", d.Reachable(7))

	fmt.Println("Path 0->5:", d.Path(5))
	fmt.Println("Order:", d.Order())
	fmt.Println("Reachable #4 #6:", d.Reachable(4), d.Reachable(6))

	// Create the breadth first search algo structure for g and source node #2
	b := bfs.New(g, 1)

	// Get the path between #0 (source) and #2
	fmt.Println("Path 0->5:", b.Path(5))
	fmt.Println("Order:", b.Order())
	fmt.Println("Reachable #4 #6:", b.Reachable(4), b.Reachable(6))
}
