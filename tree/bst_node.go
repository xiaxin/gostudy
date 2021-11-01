package tree

type BSTNode struct {
	val   *BSTValue
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(key int, val interface{}) *BSTNode {
	return &BSTNode{NewBSTValue(key, val), nil, nil}
}

func (n *BSTNode) String() string {
	return n.val.String()
}

func (n *BSTNode) Compare(other *BSTNode) int {
	return n.val.Compare(other.val)
}

func (n *BSTNode) CompareKey(key int) int {
	node := NewBSTValue(key, nil)
	return n.val.Compare(node)
}
