package common

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.PushNode(1, "A")
	stack.PushNode(2, "B")

	t.Log(stack.Pop())
	stack.PushNode(3, "C")
	t.Log(stack.Pop())
}
