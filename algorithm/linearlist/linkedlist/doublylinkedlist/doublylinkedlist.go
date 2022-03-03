package doublylinkedlist

import (
	"go-app/algorithm/linearlist"
)

// 双向链表
type DoublyLinkedList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (list *DoublyLinkedList) Size() int {
	return list.size
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoublyLinkedList) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *DoublyLinkedList) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; i < list.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *DoublyLinkedList) Add(v interface{}) {
	newNode := &Node{
		value: v,
		prev:  nil,
		next:  nil,
	}

	if list.size == 0 {
		list.first = newNode
		list.last = newNode
	} else {
		node := list.last
		node.next = newNode
		newNode.prev = node
		list.last = newNode
	}
	list.size++
}

// index从0开始, 当index为列表size时,直接添加到末尾 (index<0 || index>size返回false)
func (list *DoublyLinkedList) Insert(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		if index == list.size {
			list.Add(v)
			return true
		}
		return false
	}

	newNode := &Node{
		value: v,
		prev:  nil,
		next:  nil,
	}

	node := list.getNodeByIndex(index)
	beforeNode := node.prev

	node.prev = newNode
	newNode.next = node
	newNode.prev = beforeNode
	if beforeNode != nil {
		beforeNode.next = newNode
	} else {
		list.first = newNode
	}
	list.size++
	return true
}

func (list *DoublyLinkedList) Remove(index int) (interface{}, bool) {
	if !list.rangeCheck(index) {
		return nil, false
	}

	node := list.getNodeByIndex(index)
	val := node.value

	if node == list.first {
		list.first = node.next
	}
	if node == list.last {
		list.last = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	list.size--
	return val, true
}

func (list *DoublyLinkedList) Set(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		return false
	}

	node := list.getNodeByIndex(index)
	node.value = v
	return true
}

func (list *DoublyLinkedList) Get(index int) (interface{}, bool) {
	if !list.rangeCheck(index) {
		return nil, false
	}

	node := list.getNodeByIndex(index)
	return node.value, true
}

func (list *DoublyLinkedList) Contain(v interface{}) bool {
	return list.IndexOf(v) != -1
}

func (list *DoublyLinkedList) IndexOf(v interface{}) int {
	if list.size == 0 {
		return -1
	}

	for i, node := 0, list.first; i < list.size; i, node = i+1, node.next {
		if node.value == v {
			return i
		}
	}
	return -1
}

func (list *DoublyLinkedList) Iterator() linearlist.Iterator {
	return NewIterator(list)
}

func (list *DoublyLinkedList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}

func (list *DoublyLinkedList) getNodeByIndex(index int) *Node {
	if !list.rangeCheck(index) {
		return nil
	}

	var node *Node
	if 2*index < list.size {
		// 正向遍历
		node = list.first
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		// 反向遍历
		node = list.last
		for i := 0; i < list.size-index-1; i++ {
			node = node.prev
		}
	}
	return node
}
