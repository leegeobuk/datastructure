package list

type list[E any] interface {
	Len() int
	First() E
	Last() E
	Clear()
	Prepend(e E) bool
	Append(e E) bool
	Insert(index int, e E)
	Set(index int, e E) E
	IndexOf(e E) int
	LastIndexOf(e E) int
	Get(index int) E
	Remove(e E) bool
	RemoveAt(index int) E
	RemoveFirst() E
	RemoveLast() E
}
