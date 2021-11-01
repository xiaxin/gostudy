package common

type Stack struct {
	tail *Node
	size int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) PushNode(key int, val interface{}) {
	node := NewNode(key, val)

	s.Push(node)
}

func (s *Stack) Push(n *Node) {
	if nil == s.tail {
		s.tail = n
	} else {

		n.Prev = s.tail

		s.tail.Next = n

		s.tail = n
	}
	s.size++
}

func (s *Stack) Pop() *Node {
	node := s.tail

	if nil == node {
		return nil
	}

	s.tail = node.Prev

	if s.tail != nil {
		s.tail.Next = nil
	}

	s.size--

	return node
}

func (s *Stack) Size() int {
	return s.size
}
