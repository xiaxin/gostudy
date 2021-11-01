package tree

import "gostudy/common"

type BSTValue struct {
	*common.Node
}

func NewBSTValue(key int, val interface{}) *BSTValue {
	return &BSTValue{
		Node: common.NewNode(key, val),
	}
}

func (v *BSTValue) String() string {
	return v.Node.String()
}

func (v *BSTValue) Compare(other *BSTValue) int {
	return v.Node.Compare(other.Node)
}
