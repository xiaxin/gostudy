package tree

import (
	"reflect"
	"testing"
)

func InitBSTData() *BST {
	// 测试数据结构
	// root:4 left:2 right:6
	//   root:2 left:1 right:3
	//     root:1 left:0
	//   root:6 left:5 right:7
	bst := NewBST()
	bst.InsertNode(4, "A")
	bst.InsertNode(2, "B")
	bst.InsertNode(1, "C")
	bst.InsertNode(3, "D")
	bst.InsertNode(6, "E")
	bst.InsertNode(5, "F")
	bst.InsertNode(7, "G")
	bst.InsertNode(0, "H")
	return bst
}

func TestBSTSearch(t *testing.T) {
	// buf1 := new(bytes.Buffer)
	bst := InitBSTData()

	node := bst.Search(7)
	t.Log(node)
}

var TestBSTOrderTraverseData = []struct {
	name     string
	init     func() []string
	expected []string
}{
	{
		//  前序遍历
		"PreOrderTraverse",
		func() []string {
			return InitBSTData().PreOrderTraverse()
		},
		[]string{"[key:4] A", "[key:2] B", "[key:1] C", "[key:0] H", "[key:3] D", "[key:6] E", "[key:5] F", "[key:7] G"},
	},
	{
		//  中序遍历
		"InOrderTraverse",
		func() []string {
			return InitBSTData().InOrderTraverse()
		},
		[]string{"[key:0] H", "[key:1] C", "[key:2] B", "[key:3] D", "[key:4] A", "[key:5] F", "[key:6] E", "[key:7] G"},
	},
	{
		"PostOrderTraverse",
		func() []string {
			return InitBSTData().PostOrderTraverse()
		},
		// [key:1] C [key:3] D [key:2] B [key:5] F [key:7] G [key:6] E [key:4] A
		[]string{"[key:0] H", "[key:1] C", "[key:3] D", "[key:2] B", "[key:5] F", "[key:7] G", "[key:6] E", "[key:4] A"},
	},
}

func TestBSTOrderTraverse(t *testing.T) {

	for _, test := range TestBSTOrderTraverseData {
		t.Run(test.name, func(t *testing.T) {
			a := test.init()
			b := test.expected

			if !reflect.DeepEqual(a, b) {
				t.Errorf("Nodes should have value = %s", a)
			}
		})
	}
}

// 前序遍历
// func TestBSTPreOrderTraverse(t *testing.T) {
// 	bst := InitBSTData()

// 	a := bst.PreOrderTraverse()

// 	b := []string{"[key:4] A", "[key:2] B", "[key:1] C", "[key:0] H", "[key:3] D", "[key:6] E", "[key:5] F", "[key:7] G"}

// 	if !reflect.DeepEqual(a, b) {
// 		t.Errorf("Nodes should have value = [[key:4] A [key:2] B [key:1] C [key:0] H [key:3] D [key:6] E [key:5] F [key:7] G]")
// 	}

// }

//  中序遍历
// func TestBSTInOrderTraverse(t *testing.T) {
// 	bst := InitBSTData()

// 	inOrder := bst.InOrderTraverse()
// 	if inOrder != " [key:1] C [key:2] B [key:3] D [key:4] A [key:5] F [key:6] E [key:7] G" {
// 		t.Errorf("Unexpected %v", inOrder)
// 	}
// }

//  后续遍历
// func TestBSTPostOrderTraverse(t *testing.T) {
// 	bst := InitBSTData()

// 	postOrder := bst.PostOrderTraverse()
// 	if postOrder != " [key:1] C [key:3] D [key:2] B [key:5] F [key:7] G [key:6] E [key:4] A" {
// 		t.Errorf("Unexpected %v", postOrder)
// 	}
// }

// func TestTreeDelete(t *testing.T) {
// 	var inOrder string

// 	{
// 		bst := InitBSTData()
// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:0] H [key:1] C [key:2] B [key:3] D [key:4] A [key:5] F [key:6] E [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// 	{
// 		// 删除无子节点的节点（左）
// 		bst := InitBSTData()
// 		bst.Delete(0)
// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:1] C [key:2] B [key:3] D [key:4] A [key:5] F [key:6] E [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// 	{
// 		// 删除无子节点的节点（右）
// 		bst := InitBSTData()
// 		bst.Delete(3)
// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:0] H [key:1] C [key:2] B [key:4] A [key:5] F [key:6] E [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// 	{
// 		//  删除有左右子节点的节点 （最大节点 无 子节点）
// 		bst := InitBSTData()
// 		bst.Delete(6)

// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:0] H [key:1] C [key:2] B [key:3] D [key:4] A [key:5] F [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// 	{
// 		//  删除有左右子节点的节点 （最大节点 有 子节点）
// 		bst := InitBSTData()
// 		bst.Delete(2)

// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:0] H [key:1] C [key:3] D [key:4] A [key:5] F [key:6] E [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// 	{
// 		// 删除根
// 		bst := InitBSTData()
// 		bst.Delete(4)

// 		inOrder = bst.InOrderTraverse()
// 		if inOrder != " [key:0] H [key:1] C [key:2] B [key:3] D [key:5] F [key:6] E [key:7] G" {
// 			t.Errorf("Unexpected %v", inOrder)
// 		}
// 	}

// }
