package singlelinklist

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	list := New()
	printList(list)

	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	printList(list)
}

func printList(list *SingleLinkList) {
	fmt.Printf("list size: %d, empty: %t, value: %v\n", list.Size(), list.IsEmpty(), list.Values())
}
