package common

import "testing"

func ListData() *List {
	list := NewList()
	list.PushNode(1, "A")
	list.PushNode(2, "B")
	list.PushNode(3, "C")
	list.PushNode(4, "D")
	list.PushNode(5, "E")

	return list
}
func TestList(t *testing.T) {
	list := ListData()

	{
		node := list.Pop()
		t.Log(node)
	}

	{
		node := list.Pop()
		t.Log(node)
	}

	{
		node := list.Pop()
		t.Log(node)
	}

	{
		node := list.Pop()
		t.Log(node)
	}

	{
		node := list.Pop()
		t.Log(node)
	}

	list.PushNode(6, "F")

	{
		node := list.Pop()
		t.Log(node)
	}
}
