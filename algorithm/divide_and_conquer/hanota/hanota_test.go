package hanota

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSolveHanota(t *testing.T) {
	a := list.New()
	b := list.New()
	c := list.New()

	a.PushFront(5)
	a.PushFront(4)
	a.PushFront(3)
	a.PushFront(2)
	a.PushFront(1)
	printList(a)
	_ = b
	_ = c

	SolveHanota(a, b, c)
	printList(a)
	printList(b)
	printList(c)

	// Output:
	// list: [1 2 3 4 5]
	// list: []
	// list: []
	// list: [1 2 3 4 5]
}

func printList(l *list.List) {
	var val []interface{}
	for e := l.Front(); e != nil; e = e.Next() {
		val = append(val, e.Value)
	}
	fmt.Println("list:", val)
}
