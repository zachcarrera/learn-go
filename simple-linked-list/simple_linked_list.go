package linkedlist

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
	panic("Please implement the Push function")
}

func (l *List) Pop() (int, error) {
	panic("Please implement the Pop function")
}

func (l *List) Array() []int {
	panic("Please implement the Array function")
}

func (l *List) Reverse() *List {
	panic("Please implement the Reverse function")
}
