package bfs

import (
	graph "graph_theory/graph"
	queue "graph_theory/queue"
	stack "graph_theory/stack"
)

// Breadth-first-search algorithm data structure.
type Bfs struct {
	graph  *graph.Graph
	source int

	visited []bool
	parents []int
	order   []int
	paths   map[int][]int

	pending *queue.Queue
	builder *stack.Stack
}

// Creates a new random-order bfs structure.
func New(g *graph.Graph, src int) *Bfs {
	d := new(Bfs)

	d.graph = g
	d.source = src

	d.visited = make([]bool, g.Vertices())
	d.visited[src] = true
	d.parents = make([]int, g.Vertices())
	d.order = make([]int, 1, g.Vertices())
	d.order[0] = src
	d.paths = make(map[int][]int)

	d.pending = queue.New()
	d.pending.Push(src)
	d.builder = stack.New()

	return d
}

// Generates the path from the source node to the destination.
func (d *Bfs) Path(dst int) []int {
	// Return nil if not reachable
	if !d.Reachable(dst) {
		return nil
	}
	// If reachable, but path not yet generated, create and cache
	if cached, ok := d.paths[dst]; !ok {
		for cur := dst; cur != d.source; {
			d.builder.Push(cur)
			cur = d.parents[cur]
		}
		d.builder.Push(d.source)

		path := make([]int, d.builder.Size())
		for i := 0; i < len(path); i++ {
			path[i] = d.builder.Pop().(int)
		}
		d.paths[dst] = path
		return path
	} else {
		return cached
	}
}

// Checks whether a given vertex is reachable from the source.
func (d *Bfs) Reachable(dst int) bool {
	if !d.visited[dst] && !d.pending.Empty() {
		d.search(dst)
	}
	return d.visited[dst]
}

// Generates the full order in which nodes were traversed.
func (d *Bfs) Order() []int {
	// Force bfs termination
	if !d.pending.Empty() {
		d.search(-1)
	}
	return d.order
}

// Continues the bfs search from the last yield point, looking for dst.
func (d *Bfs) search(dst int) {
	for !d.pending.Empty() {
		// Fetch the next node, and visit if new
		src := d.pending.Pop().(int)
		d.graph.Do(src, func(peer interface{}) {
			if p := peer.(int); !d.visited[p] {
				d.visited[p] = true
				d.order = append(d.order, p)
				d.parents[p] = src
				d.pending.Push(p)
			}
		})
		// If we found the destination, yield
		if dst == src {
			return
		}
	}
}
