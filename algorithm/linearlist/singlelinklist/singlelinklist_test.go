package singlelinklist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestInsert(t *testing.T) {
	list := New()
	list.Insert(0, "b")
	list.Insert(1, "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", list.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestInsert2(t *testing.T) {
	list := New()
	for i := 0; i < 5; i++ {
		list.Insert(i, i)
	}
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, list.Values())

	list.Insert(1, 100)
	printList(list)
	list.Insert(3, 300)
	printList(list)
	list.Insert(0, -1)
	printList(list)
	list.Insert(8, 800)
	printList(list)
	list.Insert(8, 80000)
	printList(list)
	// Output:
	// list size: 6, empty: false, value: [0 100 1 2 3 4]
	// list size: 7, empty: false, value: [0 100 1 300 2 3 4]
	// list size: 8, empty: false, value: [-1 0 100 1 300 2 3 4]
	// list size: 9, empty: false, value: [-1 0 100 1 300 2 3 4 800]
	// list size: 10, empty: false, value: [-1 0 100 1 300 2 3 4 80000 800]
}

func TestRemove(t *testing.T) {
	list := New()
	val, ok := list.Remove(0)
	assert.False(t, ok)
	assert.Nil(t, val)

	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	val, ok = list.Remove(-1)
	assert.False(t, ok)
	assert.Nil(t, val)

	val, ok = list.Remove(5)
	assert.False(t, ok)
	assert.Nil(t, val)

	val, ok = list.Remove(2)
	assert.True(t, ok)
	assert.Equal(t, 2, 2)
	assert.Equal(t, []interface{}{0, 1, 3, 4}, list.Values())

	for i := 0; i < 4; i++ {
		val, ok = list.Remove(0)
		fmt.Println(val, ok)
		printList(list)
	}
}

func TestSet(t *testing.T) {
	list := New()
	assert.False(t, list.Set(0, 0))

	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	assert.False(t, list.Set(-1, -100))
	assert.False(t, list.Set(5, 500))
	printList(list)
	assert.True(t, list.Set(0, -1))  // -1 1 2 3 4
	assert.True(t, list.Set(4, 400)) // -1 1 2 3 400
	assert.True(t, list.Set(2, 200)) // -1 1 200 3 400
	printList(list)
	assert.Equal(t, []interface{}{-1, 1, 200, 3, 400}, list.Values())
}

func TestGet(t *testing.T) {
	list := New()
	val, get := list.Get(0)
	assert.False(t, get)
	assert.Nil(t, val)

	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	val, get = list.Get(-1)
	assert.False(t, get)
	assert.Nil(t, val)

	val, get = list.Get(5)
	assert.False(t, get)
	assert.Nil(t, val)

	val, get = list.Get(1)
	assert.True(t, get)
	assert.Equal(t, 1, val)
}

func TestContain(t *testing.T) {
	list := New()
	assert.False(t, list.Contain(0))

	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	assert.False(t, list.Contain(nil))
	assert.False(t, list.Contain(-1))
	assert.False(t, list.Contain(5))
	assert.True(t, list.Contain(1))

	list.Add(nil)
	assert.True(t, list.Contain(nil))
}

func TestIndexOf(t *testing.T) {
	list := New()
	assert.Equal(t, -1, list.IndexOf(0))

	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	assert.Equal(t, -1, list.IndexOf(nil))
	assert.Equal(t, -1, list.IndexOf(-1))
	assert.Equal(t, -1, list.IndexOf(5))
	assert.Equal(t, 1, list.IndexOf(1))

	list.Add(nil)
	assert.Equal(t, 5, list.IndexOf(nil))
}

func TestIterator(t *testing.T) {
	list := New()
	for i := 0; i < 5; i++ {
		list.Add(i)
	}
	printList(list)

	iterator := list.Iterator()
	var result []interface{}
	for iterator.HasNext() {
		v := iterator.Next()
		fmt.Println(v)
		result = append(result, v)
	}
	assert.Equal(t, result, list.Values())
}
