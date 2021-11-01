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
	if nil == l.tail {
		return nil
	}
	p := l.tail.Prev
	n := l.tail

	l.tail = p

	if nil != l.tail {
		l.tail.Next = nil
	} else {
		l.head = nil
	}

	l.size--

	return n
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Size() int {
	return l.size
}
