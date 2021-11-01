package graph

import "fmt"

type Edge struct {
	f *Node
	t *Node
	s float64
}

func (e *Edge) From() *Node {
	return e.f
}

func (e *Edge) To() *Node {
	return e.t
}

func (e *Edge) Score() float64 {
	return e.s
}

func (e *Edge) String() string {
	return fmt.Sprintf("[%d][%d]", e.f.ID(), e.t.ID())
}
