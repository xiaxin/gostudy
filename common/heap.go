package common

import "fmt"

// 堆
type Heap struct {
	nodes []*Node
	max   bool
}

func NewMaxHeap() *Heap {
	return &Heap{nil, true}
}

func NewMinHeap() *Heap {
	return &Heap{nil, false}
}

func (h *Heap) Size() int {
	return len(h.nodes)
}

func (h *Heap) InsertNode(key int, val interface{}) *Node {
	node := NewNode(key, val)
	return h.Insert(node)
}

func (h *Heap) Insert(node *Node) *Node {
	h.nodes = append(h.nodes, node)

	nodesLen := len(h.nodes) - 1

	h.up(nodesLen)

	return node
}

func (h *Heap) Delete() *Node {
	nodesLen := len(h.nodes)

	if nodesLen == 0 {
		return nil
	}

	node := h.nodes[0]

	h.swap(0, nodesLen-1)

	h.nodes = h.nodes[0 : nodesLen-1]

	h.down(0)

	return node
}

func (h *Heap) up(i int) {

	if i == 0 {
		return
	}

	if h.max {
		h.maxUp(i)
	} else {
		h.minUp(i)
	}
}

func (h *Heap) maxUp(i int) {
	for i > 0 {
		parentIndex := h.parent(i)

		if h.nodes[i].Compare(h.nodes[parentIndex]) == 1 {
			h.swap(i, parentIndex)
			i = parentIndex
		} else {
			break
		}
	}
}

func (h *Heap) minUp(i int) {
	for i > 0 {
		parentIndex := h.parent(i)

		if h.nodes[i].Compare(h.nodes[parentIndex]) == -1 {
			h.swap(i, parentIndex)
			i = parentIndex
		} else {
			break
		}
	}
}

func (h *Heap) down(i int) {
	if h.max {
		h.maxDown(i)
	} else {
		h.minDown(i)
	}
}

func (h *Heap) maxDown(i int) {

	nodesLen := len(h.nodes)

	for h.left(i) < nodesLen {
		maxIndex := h.left(i)
		max := h.nodes[maxIndex]

		if h.right(i) < nodesLen && max.Compare(h.nodes[h.right(i)]) == -1 {
			max = h.nodes[h.right(i)]
			maxIndex = h.right(i)
		}

		if max.Compare(h.nodes[i]) == 1 {
			h.swap(maxIndex, i)

			i = maxIndex
		} else {
			break
		}
	}
}

func (h *Heap) minDown(i int) {

	nodesLen := len(h.nodes)

	for h.left(i) < nodesLen {
		minIndex := h.left(i)
		min := h.nodes[minIndex]

		if h.right(i) < nodesLen && min.Compare(h.nodes[h.right(i)]) == 1 {
			min = h.nodes[h.right(i)]
			minIndex = h.right(i)
		}

		if min.Compare(h.nodes[i]) == -1 {
			h.swap(minIndex, i)

			i = minIndex
		} else {
			break
		}
	}
}

func (h *Heap) left(i int) int {
	return i*2 + 1
}

func (h *Heap) right(i int) int {
	return i*2 + 2
}

func (h *Heap) parent(i int) int {
	return i / 2
}

func (h *Heap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *Heap) Show(name string) {
	fmt.Printf("%s show( ", name)

	for i := 0; i < len(h.nodes); i++ {
		fmt.Printf("%s ", h.nodes[i].String())
	}

	fmt.Println(")")
}
