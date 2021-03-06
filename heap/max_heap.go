package heap

import "fmt"

type MaxHeap struct {
	items Nodes
}

func NewMax() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Show(name string) {
	fmt.Printf("%s show( ", name)

	for i := 0; i < len(h.items); i++ {
		fmt.Printf("%s ", h.items[i].String())
	}

	fmt.Println(")")
}

func (h *MaxHeap) Push(node *Node) *Node {
	h.items = append(h.items, node)

	h.swim(len(h.items) - 1)

	return node
}

func (h *MaxHeap) PushValue(val Value) *Node {
	node := NewNode(val)
	return h.Push(node)
}

func (h *MaxHeap) Pop() *Node {
	n := h.items[0]

	h.swap(0, len(h.items)-1)
	h.items = h.items[0 : len(h.items)-1]
	h.sink(0)

	return n
}

// 上浮
func (h *MaxHeap) swim(i int) {

	// parentIndex, parentNode := h.parent(i)
	// for i >= 0 && parent.Compare(h.items[i]) == -1 {
	// 	h.swap(parentIndex, i)

	// 	i = parentIndex

	// 	parentIndex, parentNode = h.parent(i)
	// }

	for i > 0 {
		parentIndex, parentNode := h.parent(i)

		if parentNode.Compare(h.items[i]) == -1 || parentNode.Compare(h.items[i]) == 0 {
			h.swap(parentIndex, i)
		}

		i = parentIndex
	}
}

// 下沉
func (h *MaxHeap) sink(i int) {
	var (
		maxNode   *Node
		rightNode *Node

		maxIndex   int
		rightIndex int
	)

	len := len(h.items)

	for h.leftIndex(i) < len {
		maxIndex, maxNode = h.left(i)

		if h.rightIndex(i) < len {
			rightIndex, rightNode = h.right(i)

			if maxNode.Compare(rightNode) == -1 {
				maxIndex = rightIndex
				maxNode = rightNode
			}
		}

		if h.items[i].Compare(maxNode) == -1 {
			h.swap(i, maxIndex)
			i = maxIndex
		} else {
			break
		}

	}
}

func (h *MaxHeap) swap(i int, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MaxHeap) parent(i int) (int, *Node) {
	index := (i - 1) / 2
	return index, h.items[index]
}

func (h *MaxHeap) left(i int) (int, *Node) {
	index := h.leftIndex(i)
	return index, h.items[index]
}

func (h *MaxHeap) leftIndex(i int) int {
	return i*2 + 1
}

func (h *MaxHeap) right(i int) (int, *Node) {
	index := h.rightIndex(i)
	return index, h.items[index]
}

func (h *MaxHeap) rightIndex(i int) int {
	return i*2 + 2
}
