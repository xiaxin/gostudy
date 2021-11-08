package backtrack

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func (n *TreeNode) ShowPreOrder(name string) {
	fmt.Printf("%s:%d\n", name, n.Val)

	if n.Left != nil {
		n.Left.ShowPreOrder("left")
	}

	if n.Right != nil {
		n.Right.ShowPreOrder("right")
	}
}

// TODO 基于 Slice 构造树
func BuildTreeBySliceInt(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := NewTreeNode(nums[0])

	BuildTreeByNode(root, 0, nums)

	return root
}

func BuildTreeByNode(root *TreeNode, index int, nums []int) {
	// 左
	leftIndex := index*2 + 1

	if leftIndex < len(nums) && nums[leftIndex] != 0 {
		root.Left = NewTreeNode(nums[leftIndex])
		BuildTreeByNode(root.Left, leftIndex, nums)
	}

	// 右
	rightIndex := index*2 + 2

	if rightIndex < len(nums) && nums[rightIndex] != 0 {
		root.Right = NewTreeNode(nums[rightIndex])
		BuildTreeByNode(root.Right, rightIndex, nums)
	}

}
