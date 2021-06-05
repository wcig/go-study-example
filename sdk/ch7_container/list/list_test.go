package list

import (
	"container/list"
	"fmt"
	"testing"
)

// container/list: 实现了双向链表

func Test(t *testing.T) {
	l := list.New()

	e4 := l.PushBack(4)
	printList(l)

	e1 := l.PushFront(1)
	printList(l)

	l.InsertBefore(3, e4)
	printList(l)

	l.InsertAfter(2, e1)
	printList(l)
	// output:
	// list: [4]
	// list: [1 4]
	// list: [1 3 4]
	// list: [1 2 3 4]
}

func printList(l *list.List) {
	var val []interface{}
	for e := l.Front(); e != nil; e = e.Next() {
		val = append(val, e.Value)
	}
	fmt.Println("list:", val)
}
