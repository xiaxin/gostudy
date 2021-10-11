package heap

import (
	"fmt"
	"testing"
)

type KeyValue struct {
	key string
	val string
}

func (kv KeyValue) String() string {
	return fmt.Sprintf("[%s]:%s", kv.key, kv.val)
}

func (kv KeyValue) Compare(v Value) int {
	if kv.val < v.(KeyValue).val {
		return -1
	} else if kv.val > v.(KeyValue).val {
		return 1
	}
	return 0
}

func Data() *Heap {

	heap := New()

	heap.PushValue(KeyValue{key: "A", val: "1"})
	heap.PushValue(KeyValue{key: "B", val: "2"})
	heap.PushValue(KeyValue{key: "C", val: "3"})
	heap.PushValue(KeyValue{key: "D", val: "4"})
	heap.PushValue(KeyValue{key: "E", val: "5"})
	heap.PushValue(KeyValue{key: "F", val: "6"})

	return heap
}

func TestPush(t *testing.T) {
	heap := Data()
	heap.Show("new")
}

func TestInsert(t *testing.T) {
	heap := Data()

	heap.Show("old")

	heap.Pop()
	heap.Show("pop 1")

	heap.Pop()
	heap.Show("pop 2")
}
