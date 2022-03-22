package list

type LinkedList[E any] struct {
	head *node[E]
	tail *node[E]
	len  int
}

type node[E any] struct {
	data E
	prev *node[E]
	next *node[E]
}

func NewLinkedList[E any]() *LinkedList[E] {
	return &LinkedList[E]{}
}

func (l *LinkedList[E]) Len() int {
	return l.len
}

func (l *LinkedList[E]) First() E {
	return l.head.data
}

func (l *LinkedList[E]) Last() E {
	return l.tail.data
}

func (l *LinkedList[E]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *LinkedList[E]) Prepend(e E) bool {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Append(e E) bool {
	n := &node[E]{data: e}
	if l.len == 0 {
		l.head = n
	} else {
		n.prev = l.tail
		l.tail.next = n
	}
	l.tail = n
	l.len++

	return true
}

func (l *LinkedList[E]) Insert(index int, e E) {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Set(index int, e E) E {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) IndexOf(e E) int {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) LastIndexOf(e E) int {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Get(index int) E {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) Remove(e E) bool {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) RemoveAt(index int) E {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) RemoveFirst() E {
	// TODO implement me
	panic("implement me")
}

func (l *LinkedList[E]) RemoveLast() E {
	// TODO implement me
	panic("implement me")
}
