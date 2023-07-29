package linkedlist

import "errors"

// Define the List and Element types here.

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
}

func New(elements []int) *List {
	baseList := &List{}
	for _, v := range elements {
		baseList.Push(v)
	}
	return baseList
}

func (l *List) Size() int {
	var size int
	current := l.head
	for current != nil {
		current = current.next
		size++
	}
	return size
}

func (l *List) Push(element int) {
	if l.head == nil {
		l.head = &Element{data: element}
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &Element{data: element}
}

func (l *List) Pop() (int, error) {

	if l.head == nil {
		return 0, errors.New("Cannot pop from an empty list")
	}

	if l.head.next == nil {
		l.head = nil
		return 0, nil
	}

	current := l.head

	for current.next.next != nil {
		current = current.next
	}

	popped := current.next.data
	current.next = nil
	return popped, nil
}

func (l *List) Array() []int {
	panic("Please implement the Array function")
}

func (l *List) Reverse() *List {
	panic("Please implement the Reverse function")
}
