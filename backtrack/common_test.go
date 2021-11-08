package backtrack

import "testing"

func TestTreeNode(t *testing.T) {
	tree := BuildTreeBySliceInt([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1})

	tree.ShowPreOrder("root")
}
