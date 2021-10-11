package heap

type Value interface {
	String() string
	Compare(val Value) int
}

type Nodes []*Node

func (n Nodes) Len() int {
	return len(n)
}

type Node struct {
	Value Value
}

func NewNode(val Value) *Node {
	return &Node{val}
}

func (n *Node) Compare(node *Node) int {
	return n.Value.Compare(node.Value)
}

func (n *Node) String() string {
	return n.Value.String()
}
