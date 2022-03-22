package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Len(t *testing.T) {
	// given
	l1 := NewLinkedList[int]()
	l2 := NewLinkedList[int]()
	l3 := NewLinkedList[int]()

	// when
	l2.Append(10)
	for i := 0; i < 100; i++ {
		l3.Append(i)
	}

	// then
	assert.Equal(t, 0, l1.Len(), "Len failed")
	assert.Equal(t, 1, l2.Len(), "Len failed")
	assert.Equal(t, 100, l3.Len(), "Len failed")
}

func TestLinkedList_First(t *testing.T) {
	// linkedList := NewLinkedList[int]()
}

func TestLinkedList_Last(t *testing.T) {
	// linkedList := NewLinkedList[int]()
}

// func TestLinkedList_Prepend(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		e E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.Prepend(tt.args.e); got != tt.want {
// 				t.Errorf("Prepend() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
func TestLinkedList_Append(t *testing.T) {
	// linkedList := NewLinkedList[int]()
}

// func TestLinkedList_Insert(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		index int
// 		e     E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			l.Insert(tt.args.index, tt.args.e)
// 		})
// 	}
// }
//
// func TestLinkedList_Set(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		index int
// 		e     E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   E
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.Set(tt.args.index, tt.args.e); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Set() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestLinkedList_IndexOf(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		e E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.IndexOf(tt.args.e); got != tt.want {
// 				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_LastIndexOf(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		e E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.LastIndexOf(tt.args.e); got != tt.want {
// 				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_Get(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		index int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   E
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_Remove(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		e E
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.Remove(tt.args.e); got != tt.want {
// 				t.Errorf("Remove() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_RemoveAt(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	type args struct {
// 		index int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   E
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.RemoveAt(tt.args.index); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RemoveAt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_RemoveFirst(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   E
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.RemoveFirst(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestLinkedList_RemoveLast(t *testing.T) {
// 	type fields struct {
// 		head *node[E]
// 		tail *node[E]
// 		size int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   E
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			l := &LinkedList{
// 				head: tt.fields.head,
// 				tail: tt.fields.tail,
// 				size: tt.fields.size,
// 			}
// 			if got := l.RemoveLast(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RemoveLast() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
