package common

import "fmt"

type Node struct {
	key int
	val interface{}

	Prev *Node
	Next *Node
}

func NewValNode(val interface{}) *Node {
	return &Node{0, val, nil, nil}
}

func NewNode(key int, val interface{}) *Node {
	return &Node{key, val, nil, nil}
}

func (n *Node) Key() int {
	return n.key
}

func (n *Node) Val() interface{} {
	return n.val
}

func (n *Node) ValString() string {
	return n.val.(string)
}

func (n *Node) IntString() int {
	return n.val.(int)
}

func (n *Node) Compare(node *Node) int {
	if n.key == node.key {
		return 0
	} else if n.key > node.key {
		return 1
	}
	return -1
}

func (n *Node) String() string {
	return fmt.Sprintf("[key:%d] %v", n.key, n.val)
}
