package heap

import "fmt"

type MinHeap struct {
	items Nodes
}

func NewMin() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Len() int {
	return h.items.Len()
}

func (h *MinHeap) PushValue(val Value) *Node {
	node := NewNode(val)
	return h.Push(node)
}

func (h *MinHeap) Push(node *Node) *Node {
	h.items = append(h.items, node)

	h.up(len(h.items) - 1)

	return node
}
func (h *MinHeap) Pop() *Node {
	n := h.items[0]

	h.swap(0, h.Len()-1)

	h.items = h.items[0 : h.Len()-1]

	h.down(0)

	return n
}

func (h *MinHeap) Show(name string) {
	fmt.Printf("%s show( ", name)

	for i := 0; i < h.Len(); i++ {
		fmt.Printf("%s", h.index(i).String())
	}

	fmt.Println(" )")
}

// 上浮
func (h *MinHeap) up(i int) {
	if i == 0 {
		return
	}
	var (
		parentIndex int
		parentNode  *Node
	)
	for i > 0 {
		parentIndex, parentNode = h.parent(i)
		if parentNode.Compare(h.index(i)) == 1 {
			h.swap(parentIndex, i)
		}
		i = parentIndex
	}
}

// 下沉
func (h *MinHeap) down(i int) {
	var (
		minIndex   int
		minNode    *Node
		rightIndex int
		rightNode  *Node
	)
	len := h.Len()

	for h.leftIndex(i) < len {
		minIndex, minNode = h.left(i)

		if h.rightIndex(i) < len {
			rightIndex, rightNode = h.right(i)

			if minNode.Compare(rightNode) == 1 {
				minIndex = rightIndex
				minNode = rightNode

			}
		}

		if minNode.Compare(h.index(i)) == -1 {
			h.swap(minIndex, i)
			i = minIndex
		} else {
			break
		}

	}
}

// 交换
func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *MinHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) leftIndex(i int) int {
	return i*2 + 1
}

func (h *MinHeap) rightIndex(i int) int {
	return i*2 + 2
}

func (h *MinHeap) parent(i int) (int, *Node) {
	index := h.parentIndex(i)
	return index, h.index(index)
}

func (h *MinHeap) left(i int) (int, *Node) {
	index := h.leftIndex(i)
	return index, h.index(index)
}

func (h *MinHeap) right(i int) (int, *Node) {
	index := h.rightIndex(i)
	return index, h.index(index)
}

func (h *MinHeap) index(i int) *Node {
	return h.items[i]
}
