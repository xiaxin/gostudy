package bst

import (
	"testing"
)

func Tr() *Tree {
	tr := New()
	tr.Insert(NodeStruct{"A", 4})
	tr.Insert(NodeStruct{"B", 2})
	tr.Insert(NodeStruct{"C", 1})
	tr.Insert(NodeStruct{"D", 3})
	tr.Insert(NodeStruct{"E", 6})
	tr.Insert(NodeStruct{"F", 5})
	tr.Insert(NodeStruct{"G", 7})
	return tr
}

func TestTreeInsert(t *testing.T) {
	// buf1 := new(bytes.Buffer)
	tr := Tr()

	preOrder := tr.PreOrder()
	if preOrder != "A(4) B(2) C(1) D(3) E(6) F(5) G(7) " {
		t.Errorf("Unexpected %v", preOrder)
	}

	inOrder := tr.InOrder()
	if inOrder != "C(1) B(2) D(3) A(4) F(5) E(6) G(7) " {
		t.Errorf("Unexpected %v", inOrder)
	}

	postOrder := tr.PostOrder()
	if postOrder != "C(1) D(3) B(2) F(5) G(7) E(6) A(4) " {
		t.Errorf("Unexpected %v", postOrder)
	}
}

func TestTreeSearch(t *testing.T) {
	// buf1 := new(bytes.Buffer)
	tr := Tr()

	node := tr.Search(NodeStruct{"A", 4})
	t.Log(node)
}

func TestTreeMinMax(t *testing.T) {
	tr := Tr()
	node := tr.Search(NodeStruct{"A", 4})

	maxNode := node.Max()
	if maxNode.String() != "G(7)" {
		t.Errorf("Unexpected %v", maxNode.String())
	}

	minNode := node.Min()
	if minNode.String() != "C(1)" {
		t.Errorf("Unexpected %v", minNode.String())
	}
}

func TrD() *Tree {
	tr := New()
	tr.Insert(NodeStruct{"A", 4})
	tr.Insert(NodeStruct{"B", 2})
	tr.Insert(NodeStruct{"C", 1})
	tr.Insert(NodeStruct{"D", 3})

	tr.Insert(NodeStruct{"E", 9})
	tr.Insert(NodeStruct{"F", 6})
	tr.Insert(NodeStruct{"G", 5})
	tr.Insert(NodeStruct{"H", 8})
	tr.Insert(NodeStruct{"I", 7})

	tr.Insert(NodeStruct{"J", 11})
	tr.Insert(NodeStruct{"K", 10})
	tr.Insert(NodeStruct{"L", 12})
	return tr
}

func TestTreeDelete(t *testing.T) {
	var inOrder string

	tr1 := TrD()
	tr1.Delete(NodeStruct{"E", 9})

	inOrder = tr1.InOrder()
	if inOrder != "C(1) B(2) D(3) A(4) G(5) F(6) I(7) H(8) K(10) J(11) L(12) " {
		t.Errorf("Unexpected %v", inOrder)
	}

	tr2 := TrD()
	tr2.Delete(NodeStruct{"I", 7})

	inOrder = tr2.InOrder()
	if inOrder != "C(1) B(2) D(3) A(4) G(5) F(6) H(8) E(9) K(10) J(11) L(12) " {
		t.Errorf("Unexpected %v", inOrder)
	}

	tr3 := TrD()
	tr3.Delete(NodeStruct{"B", 2})

	inOrder = tr3.InOrder()
	if inOrder != "C(1) D(3) A(4) G(5) F(6) I(7) H(8) E(9) K(10) J(11) L(12) " {
		t.Errorf("Unexpected %v", inOrder)
	}

	tr4 := TrD()
	tr4.Delete(NodeStruct{"J", 11})

	inOrder = tr4.InOrder()
	if inOrder != "C(1) B(2) D(3) A(4) G(5) F(6) I(7) H(8) E(9) K(10) L(12) " {
		t.Errorf("Unexpected %v", inOrder)
	}

	tr5 := TrD()
	tr5.Delete(NodeStruct{"A", 4})

	inOrder = tr5.InOrder()
	if inOrder != "C(1) B(2) D(3) A(4) G(5) F(6) I(7) H(8) E(9) K(10) L(12) " {
		t.Errorf("Unexpected %v", inOrder)
	}

}
