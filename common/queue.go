package common

type PrioQueue struct {
	heap *Heap
}

func NewPrioQueue() *PrioQueue {
	return &PrioQueue{
		heap: NewMaxHeap(),
	}
}

func (q *PrioQueue) Push(key int, val interface{}) {
	q.heap.InsertNode(key, val)
}

func (q *PrioQueue) Pop() *Node {
	return q.heap.Delete()
}
