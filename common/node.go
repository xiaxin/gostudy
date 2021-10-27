package common

import "fmt"

type Node struct {
	key int
	val interface{}
}

func NewNode(key int, val interface{}) *Node {
	return &Node{key, val}
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
