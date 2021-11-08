package common

type List struct {
	head *Node
	tail *Node

	size int
}

func NewList() *List {
	return &List{}
}

func (l *List) PushNode(key int, val interface{}) {
	node := NewNode(key, val)

	l.Push(node)
}

func (l *List) Push(n *Node) {
	if nil == l.head {
		l.head = n
		l.tail = n
	} else {
		n.Prev = l.tail
		l.tail.Next = n
		l.tail = n
	}

	l.size++
}

func (l *List) Pop() *Node {
	if nil == l.head {
		return nil
	}

	h := l.head

	l.head = h.Next

	if l.head == nil {
		l.tail = nil
	}

	l.size--

	return h
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Size() int {
	return l.size
}
