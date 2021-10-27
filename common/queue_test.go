package common

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(1, "A")
	queue.Push(2, "B")

	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
}
