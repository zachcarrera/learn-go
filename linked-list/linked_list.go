package linkedlist

// Define List and Node types here.
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
	tail *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	panic("Please implement the NewList function")
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	if l.head == nil {
		l.Push(v)
		return
	}

	l.head.prev = &Node{Value: v, next: l.head}
	l.head = l.head.prev
}

func (l *List) Push(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
		l.tail = l.head
		return
	}

	l.tail.next = &Node{Value: v, prev: l.tail}
	l.tail = l.tail.next
}

func (l *List) Shift() (interface{}, error) {
	panic("Please implement the Shift function")
}

func (l *List) Pop() (interface{}, error) {
	panic("Please implement the Pop function")
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
