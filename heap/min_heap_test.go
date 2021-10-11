package heap

import "testing"

func MinData() *MinHeap {
	heap := NewMin()

	heap.PushValue(KeyValue{key: "A", val: "1"})
	heap.PushValue(KeyValue{key: "B", val: "2"})
	heap.PushValue(KeyValue{key: "C", val: "3"})
	heap.PushValue(KeyValue{key: "D", val: "4"})
	heap.PushValue(KeyValue{key: "E", val: "5"})
	heap.PushValue(KeyValue{key: "F", val: "6"})
	heap.PushValue(KeyValue{key: "G", val: "7"})

	return heap
}

func TestMinPush(t *testing.T) {
	heap := MinData()
	heap.Show("new")
}

func TestMinInsert(t *testing.T) {
	heap := MinData()

	heap.Show("old")

	heap.Pop()
	heap.Show("pop 1")

	heap.Pop()
	heap.Show("pop 2")
}
