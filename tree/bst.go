package tree

import "fmt"

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{nil}
}

//  判断是否为空
func (t *BST) IsEmpty() bool {
	return t.root == nil
}

// 搜索
func (t *BST) Search(key int) *BSTNode {
	_, node := t.SearchParent(key)
	return node
}

// 搜索
func (t *BST) SearchParent(key int) (*BSTNode, *BSTNode) {
	var parentNode *BSTNode
	node := t.root

	for node != nil {
		if node.CompareKey(key) == 0 {
			return parentNode, node
		} else if node.CompareKey(key) == 1 {
			parentNode = node
			node = node.Left
		} else if node.CompareKey(key) == -1 {
			parentNode = node
			node = node.Right
		}
	}

	return nil, nil
}

// 插入节点
func (t *BST) InsertNode(key int, val interface{}) *BSTNode {
	node := NewBSTNode(key, val)
	return t.Insert(node)
}

// 插入
func (t *BST) Insert(node *BSTNode) *BSTNode {
	if nil == t.root {
		t.root = node
	} else {
		root := t.root
		for root != nil {
			if node.Compare(root) == -1 {
				if root.Left == nil {
					root.Left = node
					return node
				}
				root = root.Left
			} else if node.Compare(root) == 1 || node.Compare(root) == 0 {
				if root.Right == nil {
					root.Right = node
					return node
				}
				root = root.Right
			}
		}
	}

	return node
}

// 删除 TODO
func (t *BST) Delete(key int) *BSTNode {
	parentNode, node := t.SearchParent(key)

	// 没找到
	if node == nil {
		return nil
	}

	// //  根
	// if parentNode == nil && node != nil {
	// 	t.root = nil
	// 	return node
	// }

	// 普通情况
	if node.Left == nil && node.Right == nil {

		if nil == parentNode {
			t.root = nil
		} else {
			// 左右节点为空
			if parentNode.Compare(node) == -1 || parentNode.Compare(node) == 0 {
				parentNode.Right = nil
			} else if parentNode.Compare(node) == 1 {
				parentNode.Left = nil
			}
		}
	} else if node.Left != nil && node.Right == nil {
		// 左节点不为空
		if nil == parentNode {
			t.root = node.Left
		} else {
			if parentNode.Compare(node) == -1 || parentNode.Compare(node) == 0 {
				parentNode.Right = node.Left
			} else if parentNode.Compare(node) == 1 {
				parentNode.Left = node.Left
			}
		}

	} else if node.Left == nil && node.Right != nil {
		if nil == parentNode {
			t.root = node.Right
		} else {
			// 右节点不为空
			if parentNode.Compare(node) == -1 || parentNode.Compare(node) == 0 {
				parentNode.Right = node.Right
			} else if parentNode.Compare(node) == 1 {
				parentNode.Left = node.Right
			}
		}

	} else if node.Left != nil && node.Right != nil {
		// 都不为空 TODO 备注
		maxParentNode, maxNode := t.MaxParent(node.Left)

		// 边界判断逻辑
		// 1. 判断删除节点是否为根节点（parentNode == nil）
		// 2. 判断最大节点的父节点不是删除节点（删除节点的左右节点不用设置）
		// 3. 判断删除节点的左节点不是最大节点（左节点如果是最大节点，做大节点的左子节点可能会指向自己造成死循环）
		// 4. 判断删除节点的右节点不是最大节点（右节点如果是最大节点，做大节点的右子节点可能会指向自己造成死循环）
		if parentNode == nil {
			t.root = maxNode
		} else {
			if parentNode.Compare(maxNode) == -1 || parentNode.Compare(maxNode) == 0 {
				parentNode.Right = maxNode
			} else if parentNode.Compare(maxNode) == 1 {
				parentNode.Left = maxNode
			}
		}

		if maxParentNode != node && maxParentNode != nil {
			// 最大节点的左子节点一定比最大节点的父节点大
			maxParentNode.Right = maxNode.Left
		}

		if node.Left != maxNode {
			maxNode.Left = node.Left
		}

		if node.Right != maxNode {
			maxNode.Right = node.Right
		}

	}

	return nil
}

// 前序
func (t *BST) PreOrderTraverse() string {
	return t.preOrderTraverse(t.root)
}

func (t *BST) preOrderTraverse(node *BSTNode) string {
	if nil == node {
		return ""
	}

	result := fmt.Sprintf(" %s", node.String())

	result += t.preOrderTraverse(node.Left)
	result += t.preOrderTraverse(node.Right)

	return result
}

// 中序
func (t *BST) InOrderTraverse() string {
	return t.inOrderTraverse(t.root)
}

func (t *BST) inOrderTraverse(node *BSTNode) string {
	if nil == node {
		return ""
	}

	result := t.inOrderTraverse(node.Left)
	result += fmt.Sprintf(" %s", node.String())
	result += t.inOrderTraverse(node.Right)

	return result
}

// 后续
func (t *BST) PostOrderTraverse() string {
	return t.postOrderTraverse(t.root)
}

func (t *BST) postOrderTraverse(node *BSTNode) string {
	if nil == node {
		return ""
	}

	result := t.postOrderTraverse(node.Left)
	result += t.postOrderTraverse(node.Right)
	result += fmt.Sprintf(" %s", node.String())

	return result
}

func (t *BST) Max(n *BSTNode) *BSTNode {
	_, node := t.MaxParent(n)
	return node
}

func (t *BST) Min(n *BSTNode) *BSTNode {
	_, node := t.MinParent(n)
	return node
}

func (t *BST) MinParent(node *BSTNode) (*BSTNode, *BSTNode) {
	if nil == node {
		return nil, nil
	}
	var parentNode *BSTNode

	for node.Left != nil {
		parentNode = node
		node = node.Left
	}

	return parentNode, node
}

func (t *BST) MaxParent(node *BSTNode) (*BSTNode, *BSTNode) {
	if node == nil {
		return nil, nil
	}

	var parentNode *BSTNode

	for node.Right != nil {
		parentNode = node
		node = node.Right
	}

	return parentNode, node
}
