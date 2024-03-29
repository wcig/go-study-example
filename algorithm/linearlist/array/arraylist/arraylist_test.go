package arraylist

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	list := new(ArrayList)
	printList(list)

	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	printList(list)

	insertResult := list.Insert(5, 50)
	fmt.Println("insert result:", insertResult)
	printList(list)

	removeValue, removeResult := list.Remove(5)
	fmt.Printf("remove value: %v, result: %t\n", removeValue, removeResult)
	printList(list)

	setResult := list.Set(5, 500)
	fmt.Println("set result:", setResult)
	printList(list)

	getValue, getResult := list.Get(5)
	fmt.Printf("get value: %v, result: %t\n", getValue, getResult)
	printList(list)

	containResult := list.Contain(getValue)
	fmt.Println("contain result:", containResult)
	printList(list)

	indexValue := list.IndexOf(getValue)
	fmt.Println("index value:", indexValue)
	printList(list)

	values := list.Values()
	fmt.Println("values:", values)
	printList(list)

	list.Clear()
	printList(list)
	// Output:
	// list size: 0, empty: true, value: []
	// list size: 10, empty: false, value: [0 1 2 3 4 5 6 7 8 9]
	// insert result: true
	// list size: 11, empty: false, value: [0 1 2 3 4 50 5 6 7 8 9]
	// remove value: 50, result: true
	// list size: 10, empty: false, value: [0 1 2 3 4 5 6 7 8 9]
	// set result: true
	// list size: 10, empty: false, value: [0 1 2 3 4 500 6 7 8 9]
	// get value: 500, result: true
	// list size: 10, empty: false, value: [0 1 2 3 4 500 6 7 8 9]
	// contain result: true
	// list size: 10, empty: false, value: [0 1 2 3 4 500 6 7 8 9]
	// index value: 5
	// list size: 10, empty: false, value: [0 1 2 3 4 500 6 7 8 9]
	// values: [0 1 2 3 4 500 6 7 8 9]
	// list size: 10, empty: false, value: [0 1 2 3 4 500 6 7 8 9]
	// list size: 0, empty: true, value: []
}

func printList(list *ArrayList) {
	fmt.Printf("list size: %d, empty: %t, value: %v\n", list.Size(), list.IsEmpty(), list.Values())
}

func TestAdd(t *testing.T) {
	list := New()
	for i := 0; i < 20; i++ {
		list.Add(i)
		fmt.Printf("size: %d, capacity: %d, list: %v\n", len(list.data), cap(list.data), list.Values())
	}
	// Output:
	// size: 1, capacity: 10, list: [0]
	// 	size: 2, capacity: 10, list: [0 1]
	// 	size: 3, capacity: 10, list: [0 1 2]
	// 	size: 4, capacity: 10, list: [0 1 2 3]
	// 	size: 5, capacity: 10, list: [0 1 2 3 4]
	// 	size: 6, capacity: 10, list: [0 1 2 3 4 5]
	// 	size: 7, capacity: 10, list: [0 1 2 3 4 5 6]
	// 	size: 8, capacity: 10, list: [0 1 2 3 4 5 6 7]
	// 	size: 9, capacity: 10, list: [0 1 2 3 4 5 6 7 8]
	// 	size: 10, capacity: 10, list: [0 1 2 3 4 5 6 7 8 9]
	// 	size: 11, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10]
	// 	size: 12, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11]
	// 	size: 13, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12]
	// 	size: 14, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
	// 	size: 15, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	// 	size: 16, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	// 	size: 17, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	// 	size: 18, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]
	// 	size: 19, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18]
	// 	size: 20, capacity: 20, list: [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
}

func TestInsert(t *testing.T) {
	list := New()
	for i := 0; i < 10; i++ {
		result := list.Insert(i, i)
		fmt.Printf("result: %t, size: %d, capacity: %d, list: %v\n", result, len(list.data), cap(list.data), list.Values())
	}
	// Output:
	// result: true, size: 1, capacity: 10, list: [0]
	// result: true, size: 2, capacity: 10, list: [0 1]
	// result: true, size: 3, capacity: 10, list: [0 1 2]
	// result: true, size: 4, capacity: 10, list: [0 1 2 3]
	// result: true, size: 5, capacity: 10, list: [0 1 2 3 4]
	// result: true, size: 6, capacity: 10, list: [0 1 2 3 4 5]
	// result: true, size: 7, capacity: 10, list: [0 1 2 3 4 5 6]
	// result: true, size: 8, capacity: 10, list: [0 1 2 3 4 5 6 7]
	// result: true, size: 9, capacity: 10, list: [0 1 2 3 4 5 6 7 8]
	// result: true, size: 10, capacity: 10, list: [0 1 2 3 4 5 6 7 8 9]
}

func TestRemove(t *testing.T) {
	list := New()
	for i := 0; i < 20; i++ {
		list.Add(i)
	}
	fmt.Println(list.Values())

	for i := 0; i < 20; i++ {
		element, result := list.Remove(0)
		fmt.Printf("element: %v, result: %t, size: %d, capacity: %d, list: %v\n", element, result, len(list.data), cap(list.data), list.Values())
	}
	// Output:
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 0, result: true, size: 19, capacity: 20, list: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 1, result: true, size: 18, capacity: 20, list: [2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 2, result: true, size: 17, capacity: 20, list: [3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 3, result: true, size: 16, capacity: 20, list: [4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 4, result: true, size: 15, capacity: 20, list: [5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 5, result: true, size: 14, capacity: 20, list: [6 7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 6, result: true, size: 13, capacity: 20, list: [7 8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 7, result: true, size: 12, capacity: 20, list: [8 9 10 11 12 13 14 15 16 17 18 19]
	// element: 8, result: true, size: 11, capacity: 20, list: [9 10 11 12 13 14 15 16 17 18 19]
	// element: 9, result: true, size: 10, capacity: 20, list: [10 11 12 13 14 15 16 17 18 19]
	// element: 10, result: true, size: 9, capacity: 20, list: [11 12 13 14 15 16 17 18 19]
	// element: 11, result: true, size: 8, capacity: 20, list: [12 13 14 15 16 17 18 19]
	// element: 12, result: true, size: 7, capacity: 20, list: [13 14 15 16 17 18 19]
	// element: 13, result: true, size: 6, capacity: 20, list: [14 15 16 17 18 19]
	// element: 14, result: true, size: 5, capacity: 5, list: [15 16 17 18 19]
	// element: 15, result: true, size: 4, capacity: 5, list: [16 17 18 19]
	// element: 16, result: true, size: 3, capacity: 5, list: [17 18 19]
	// element: 17, result: true, size: 2, capacity: 5, list: [18 19]
	// element: 18, result: true, size: 1, capacity: 1, list: [19]
	// element: 19, result: true, size: 0, capacity: 0, list: []
}

func TestIterator(t *testing.T) {
	list := New()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	printList(list)

	iterator := list.Iterator()
	for iterator.HasNext() {
		val := iterator.Next()
		fmt.Printf("%v, ", val)
	}
	fmt.Println()
	fmt.Println(iterator.HasNext(), iterator.Next())
	// Output:
	// list size: 10, empty: false, value: [0 1 2 3 4 5 6 7 8 9]
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	// false <nil>
}
