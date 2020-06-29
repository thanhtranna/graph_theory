package dfs_test

import (
	"fmt"

	"graph_theory/graph"
	"graph_theory/graph/dfs"
)

// Small API demo based on a trie graph and a few disconnected vertices.
func Example_usage() {
	// Create the graph
	g := graph.New(7)
	g.Connect(0, 1)
	g.Connect(1, 2)
	g.Connect(2, 3)
	g.Connect(3, 4)
	g.Connect(3, 5)

	// Create the depth first search algo structure for g and source node #2
	d := dfs.New(g, 0)

	// Get the path between #0 (source) and #2
	fmt.Println("Path 0->5:", d.Path(5))
	fmt.Println("Order:", d.Order())
	fmt.Println("Reachable #4 #6:", d.Reachable(4), d.Reachable(6))

	// Output:
	// Path 0->5: [0 1 2 3 5]
	// Order: [0 1 2 3 5 4]
	// Reachable #4 #6: true false
}
