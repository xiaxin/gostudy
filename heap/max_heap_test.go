package heap

import (
	"testing"
)

func MaxData() *MaxHeap {

	heap := NewMax()

	heap.PushValue(KeyValue{key: "A", val: "1"})
	heap.PushValue(KeyValue{key: "B", val: "2"})
	heap.PushValue(KeyValue{key: "C", val: "3"})
	heap.PushValue(KeyValue{key: "D", val: "4"})
	heap.PushValue(KeyValue{key: "E", val: "5"})
	heap.PushValue(KeyValue{key: "F", val: "6"})

	return heap
}

func TestMaxPush(t *testing.T) {
	heap := MaxData()
	heap.Show("new")
}

func TestMaxInsert(t *testing.T) {
	heap := MaxData()

	heap.Show("old")

	heap.Pop()
	heap.Show("pop 1")

	heap.Pop()
	heap.Show("pop 2")
}
