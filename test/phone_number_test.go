package test

import (
	"fmt"
	"testing"
)

type Node struct {
	key int
	val interface{}

	Prev *Node
	Next *Node
}

func NewValNode(val interface{}) *Node {
	return &Node{0, val, nil, nil}
}

func NewNode(key int, val interface{}) *Node {
	return &Node{key, val, nil, nil}
}

func (n *Node) Key() int {
	return n.key
}

func (n *Node) Val() interface{} {
	return n.val
}

func (n *Node) ValString() string {
	return n.val.(string)
}

func (n *Node) IntString() int {
	return n.val.(int)
}

func (n *Node) Compare(node *Node) int {
	if n.key == node.key {
		return 0
	} else if n.key > node.key {
		return 1
	}
	return -1
}

func (n *Node) String() string {
	return fmt.Sprintf("[key:%d] %v", n.key, n.val)
}

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

func letterCombinations(input string) []string {
	keymap := map[rune][]rune{
		'1': {},
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	queue := NewList()

	inputRune := []rune(input)

	for i := 0; i < len(inputRune); i++ {
		if inputRune[i] == '1' {
			continue
		}

		phoneNumber(keymap[inputRune[i]], queue)
	}

	var result []string
	for queue.Size() > 0 {
		node := queue.Pop()

		result = append(result, node.ValString())
	}

	return result
}

func phoneNumber(list []rune, queue *List) []string {
	var (
		result []string
		bases  []string
	)

	for queue.Size() > 0 {
		node := queue.Pop()
		bases = append(bases, node.ValString())

	}

	if len(bases) > 0 {
		for _, base := range bases {
			for _, v := range list {
				queue.PushNode(0, base+string(v))
			}
		}
	} else {
		for _, v := range list {
			queue.PushNode(0, string(v))
		}
	}

	return result
}

func TestPhoneNumber(t *testing.T) {
	list := letterCombinations("2")

	t.Log(list)
}
