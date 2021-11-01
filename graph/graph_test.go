package graph

import "testing"

func GraphData() *Graph {
	graph := NewGraph()

	// 节点
	n1 := graph.AddNode(&Node{1, nil})
	n2 := graph.AddNode(&Node{2, nil})
	n3 := graph.AddNode(&Node{3, nil})
	n4 := graph.AddNode(&Node{4, nil})
	n5 := graph.AddNode(&Node{5, nil})

	// n6
	graph.AddNode(&Node{6, nil})

	// 边
	graph.SetEdge(n1, n2, 0)
	graph.SetEdge(n1, n3, 0)

	graph.SetEdge(n2, n3, 0)
	graph.SetEdge(n2, n4, 0)

	graph.SetEdge(n3, n4, 0)
	graph.SetEdge(n3, n5, 0)

	return graph
}

func TestGraph(t *testing.T) {
	graph := GraphData()

	t.Log(graph)
}

func TestGraphBFS(t *testing.T) {

	graph := GraphData()

	res := graph.BFS(1, 6)

	t.Log(res)
}

func TestGraphDFS(t *testing.T) {

	graph := GraphData()

	res := graph.DFS(1, 6)

	t.Log(res)
}
