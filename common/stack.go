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

func (s *Stack) String() string {
	var result = ""
	node := s.tail
	for {

		if nil == node {
			break
		}
		result = node.String() + result

		node = node.Prev
	}

	return result
}

func (s *Stack) ValString() string {
	var result = ""
	node := s.tail
	for {

		if nil == node {
			break
		}
		result = node.ValString() + result

		node = node.Prev
	}

	return result
}
