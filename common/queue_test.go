package common

import "testing"

func TestPrioQueue(t *testing.T) {
	queue := NewPrioQueue()
	queue.Push(1, "A")
	queue.Push(2, "B")

	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
}
