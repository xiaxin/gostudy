package bst

import "fmt"

// 二叉查找树

// 树
type Tree struct {
	Root *Node
}

// 节点
type Node struct {
	Value Value
	Left  *Node
	Right *Node
}

// 值接口
type Value interface {
	/**
	  = 0
	  < -1
	  > 1
	**/
	Compare(value Value) int
	String() string
}

func New() *Tree {
	return &Tree{nil}
}

func (t *Tree) IsEmpty() bool {
	return t.Root == nil
}

// 搜索
func (t *Tree) Search(val Value) *Node {
	node := NewNode(val)
	return t.Root.Search(node)
}

// 插入
func (t *Tree) Insert(val Value) {
	node := NewNode(val)
	if t.Root == nil {
		t.Root = node
	} else {
		t.Root.Insert(node)
	}
}

// 删除
func (t *Tree) Delete(val Value) *Node {
	parentNode, currentNode := t.SearchParent(val)

	if nil == currentNode {
		return nil
	}

	if currentNode.Left != nil && currentNode.Right != nil {
		// 左右子节点不为空
		maxParentNode, maxNode := currentNode.Left.MaxParent()

		maxLeftNode := maxNode.Left

		if parentNode.Compare(maxNode) == 0 || parentNode.Compare(maxNode) == -1 {
			parentNode.Right = maxNode

		} else {
			parentNode.Left = maxNode
		}

		// 最大节点不能是当前节点的右节点
		if maxNode != currentNode.Right {
			maxNode.Right = currentNode.Right
		}

		// 最大节点不能是当前节点的左节点
		if maxNode != currentNode.Left {
			maxNode.Left = currentNode.Left
		}

		if nil != maxParentNode {
			maxParentNode.Right = maxLeftNode
		}

		// maxNode.Insert(maxLeftNode)

	} else if currentNode.Left != nil && currentNode.Right == nil {
		// 左子节点不为空
		if parentNode == nil {
			t.Root = currentNode.Left
		} else {
			val := parentNode.Compare(currentNode)
			switch val {
			case 0, -1:
				parentNode.Right = currentNode.Left
			case 1:
				parentNode.Left = currentNode.Left
			}
		}
	} else if currentNode.Left == nil && currentNode.Right != nil {
		if parentNode == nil {
			t.Root = currentNode.Right
		} else {
			val := parentNode.Compare(currentNode)
			switch val {
			case 0, -1:
				parentNode.Right = currentNode.Right
			case 1:
				parentNode.Left = currentNode.Right
			}
		}

	} else if currentNode.Left == nil && currentNode.Right == nil {
		// 叶子节点，把父节点的关联关系解除
		if parentNode == nil {
			t.Root = nil
		} else {
			val := parentNode.Compare(currentNode)

			switch val {
			case 0, -1:
				parentNode.Right = nil
			case 1:
				parentNode.Left = nil
			}
		}
	}

	return currentNode
}

func (t *Tree) SearchParent(val Value) (*Node, *Node) {
	node := NewNode(val)
	return t.Root.SearchParent(node)
}

func (t *Tree) PreOrder() string {
	return t.preOrder(t.Root)
}

func (t *Tree) preOrder(node *Node) string {
	if node == nil {
		return ""
	}

	res := fmt.Sprintf("%s ", node.String())

	res += t.preOrder(node.Left)
	res += t.preOrder(node.Right)
	return res
}

func (t *Tree) InOrder() string {
	return t.inOrder(t.Root)
}

func (t *Tree) inOrder(node *Node) string {
	if node == nil {
		return ""
	}

	res := t.inOrder(node.Left)
	res += fmt.Sprintf("%s ", node.String())
	res += t.inOrder(node.Right)
	return res
}

func (t *Tree) PostOrder() string {
	return t.postOrder(t.Root)
}

func (t *Tree) postOrder(node *Node) string {
	if node == nil {
		return ""
	}

	res := t.postOrder(node.Left)
	res += t.postOrder(node.Right)
	res += fmt.Sprintf("%s ", node.String())
	return res
}

func NewNode(val Value) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}

// TODO
func (n *Node) Insert(node *Node) *Node {
	if n == nil {
		return node
	}

	if n.Compare(node) == 0 || n.Compare(node) == 1 {
		n.Left = n.Left.Insert(node)
	} else {
		n.Right = n.Right.Insert(node)
	}
	return n
}

func (n *Node) Search(node *Node) *Node {

	val := node.Compare(node)

	switch val {
	case 0:
		return n
	case 1:
		return node.Left.Search(node)
	case -1:
		return node.Right.Search(node)
	}

	return nil
}

//  搜索父元素
func (n *Node) SearchParent(node *Node) (*Node, *Node) {
	var (
		parent  *Node
		current *Node = n
	)

	for nil != current {
		val := current.Compare(node)

		switch val {
		case 0:
			return parent, current
		case 1:
			parent = current
			current = current.Left
		case -1:
			parent = current
			current = current.Right

		}
	}
	return nil, current
}

func (n *Node) Compare(node *Node) int {
	return n.Value.Compare(node.Value)
}

// 返回 最小元素
func (n *Node) Min() *Node {
	_, min := n.MinParent()
	return min
}

func (n *Node) MinParent() (*Node, *Node) {
	node := n

	if nil == node {
		return nil, nil
	}

	var parentNode *Node
	for node.Left != nil {
		parentNode = node
		node = node.Left
	}

	return parentNode, node
}

// 返回最大元素
func (n *Node) Max() *Node {
	_, max := n.MaxParent()
	return max
}

func (n *Node) MaxParent() (*Node, *Node) {
	node := n

	if nil == node {
		return nil, nil
	}
	var parentNode *Node

	for node.Right != nil {
		parentNode = node
		node = node.Right
	}

	return parentNode, node
}

func (n *Node) String() string {
	return n.Value.String()
}

type NodeStruct struct {
	Key   string
	Value int
}

func (n NodeStruct) Compare(b Value) int {
	if n.Value < b.(NodeStruct).Value {
		return -1
	} else if n.Value > b.(NodeStruct).Value {
		return 1
	}
	return 0
}

func (n NodeStruct) String() string {
	return fmt.Sprintf("%s(%d)", n.Key, n.Value)
}
