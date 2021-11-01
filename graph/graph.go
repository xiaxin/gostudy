package graph

import "gostudy/common"

type Graph struct {
	nodes map[int]*Node
	edges map[int]map[int]*Edge
}

func NewGraph() *Graph {
	graph := &Graph{
		nodes: make(map[int]*Node),
		edges: make(map[int]map[int]*Edge),
	}

	return graph
}

func (g *Graph) AddNode(n *Node) *Node {
	nid := n.ID()

	g.nodes[nid] = n
	g.edges[nid] = make(map[int]*Edge)

	return n
}

func (g *Graph) DelNode(n *Node) {
	nid := n.ID()

	delete(g.nodes, nid)

	for tid := range g.edges[nid] {
		delete(g.edges[tid], nid)
	}

	delete(g.edges, nid)
}

func (g *Graph) SetEdge(f, t *Node, s float64) {
	if _, ok := g.nodes[f.ID()]; !ok {
		g.AddNode(f)
	}

	if _, ok := g.nodes[t.ID()]; !ok {
		g.AddNode(t)
	}

	edge := &Edge{f, t, s}

	g.edges[f.ID()][t.ID()] = edge
	g.edges[t.ID()][f.ID()] = edge
}

func (g *Graph) DelEdge(e *Edge) {
	f, t := e.From(), e.To()

	if _, ok := g.nodes[f.ID()]; !ok {
		return
	}

	if _, ok := g.nodes[t.ID()]; !ok {
		return
	}

	delete(g.edges[f.ID()], t.ID())
	delete(g.edges[t.ID()], f.ID())
}

func (g *Graph) Has(n *Node) bool {
	_, ok := g.nodes[n.ID()]
	return ok
}

func (g *Graph) Nodes() {

}

func (g *Graph) Edges() {

}

func (g *Graph) BFS(f, t int) bool {
	queue := common.NewList()

	visited := make(map[string]bool)

	if _, ok := g.nodes[t]; !ok {
		return false
	}

	if node, ok := g.nodes[f]; ok {
		queue.PushNode(node.ID(), nil)

		for queue.Size() > 0 {
			node := queue.Pop()

			edges := g.edges[node.Key()]

			for _, edge := range edges {

				if _, ok := visited[edge.String()]; ok {
					continue
				}

				if edge.To().ID() == t {
					return true
				}

				queue.PushNode(edge.To().ID(), nil)

				visited[edge.String()] = true
			}
		}
	}
	return false
}

func (g *Graph) DFS(f, t int) bool {
	queue := common.NewStack()

	visited := make(map[string]bool)

	if node, ok := g.nodes[f]; ok {
		queue.PushNode(node.ID(), nil)

		for queue.Size() > 0 {
			node := queue.Pop()

			edges := g.edges[node.Key()]

			for _, edge := range edges {

				if _, ok := visited[edge.String()]; ok {
					continue
				}

				if edge.To().ID() == t {
					return true
				}

				queue.PushNode(edge.To().ID(), nil)

				visited[edge.String()] = true
			}
		}
	}
	return false
}
