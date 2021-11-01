package graph

type Node struct {
	id  int
	val interface{}
}

func (n *Node) ID() int {
	return n.id
}

func (n *Node) Val() interface{} {
	return n.val
}
