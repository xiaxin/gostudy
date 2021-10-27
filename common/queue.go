package common

type Queue struct {
	heap *Heap
}

func NewQueue() *Queue {
	return &Queue{
		heap: NewMaxHeap(),
	}
}

func (q *Queue) Push(key int, val interface{}) {
	q.heap.InsertNode(key, val)
}

func (q *Queue) Pop() *Node {
	return q.heap.Delete()
}
