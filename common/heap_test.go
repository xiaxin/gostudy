package common

import (
	"testing"
)

func MaxData() *Heap {
	heap := NewMaxHeap()
	return InsertData(heap)
}

func MinData() *Heap {
	heap := NewMinHeap()
	return InsertData(heap)
}

func InsertData(heap *Heap) *Heap {
	heap.InsertNode(1, "A")
	heap.InsertNode(2, "B")
	heap.InsertNode(3, "C")
	heap.InsertNode(4, "D")
	heap.InsertNode(5, "E")
	heap.InsertNode(6, "F")

	return heap
}

func TestMaxData(t *testing.T) {
	heap := MaxData()
	heap.Show("new")
}

func TestMinData(t *testing.T) {
	heap := MinData()
	heap.Show("new")
}

func TestMaxDelete(t *testing.T) {
	heap := MaxData()

	heap.Show("old")

	heap.Delete()
	heap.Show("pop 1")

	heap.Delete()
	heap.Show("pop 2")
}

func TestMinDelete(t *testing.T) {
	heap := MinData()

	heap.Show("old")

	heap.Delete()
	heap.Show("pop 1")

	heap.Delete()
	heap.Show("pop 2")
}
