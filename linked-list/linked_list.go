package linkedlist

import "errors"

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
	newList := &List{}
	for _, v := range elements {
		newList.Push(v)
	}
	return newList
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
	if l.head == nil {
		return nil, errors.New("Cannot shift an empty list")
	}
	if l.head == l.tail {
		poppedValue := l.head.Value
		l.head, l.tail = nil, nil
		return poppedValue, nil
	}
	poppedNode := l.head
	l.head = l.head.next
	l.head.prev = nil
	return poppedNode.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("Cannot pop from an empty list")
	}
	if l.head == l.tail {
		poppedValue := l.tail.Value
		l.head, l.tail = nil, nil
		return poppedValue, nil
	}
	poppedNode := l.tail
	l.tail = poppedNode.prev
	poppedNode.prev = nil
	l.tail.next = nil
	return poppedNode.Value, nil
}

func (l *List) Reverse() {
	l.head, l.tail = l.tail, l.head
	current := l.head
	for current != nil {
		current.next, current.prev = current.prev, current.next
		current = current.next
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
