package doublycircularlinkedlist

import "go-app/algorithm/linearlist"

// 双向循环链表
type DoubleCircularLinkedList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func New() *DoubleCircularLinkedList {
	return &DoubleCircularLinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (list *DoubleCircularLinkedList) Size() int {
	return list.size
}

func (list *DoubleCircularLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoubleCircularLinkedList) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *DoubleCircularLinkedList) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; i < list.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *DoubleCircularLinkedList) ValuesReverse() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.last; i < list.size; i, node = i+1, node.prev {
		values[i] = node.value
	}
	return values
}

func (list *DoubleCircularLinkedList) Add(v interface{}) {
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
	list.last.next = list.first
	list.first.prev = list.last
	list.size++
}

// index从0开始, 当index为列表size时,直接添加到末尾 (index<0 || index>size返回false)
func (list *DoubleCircularLinkedList) Insert(index int, v interface{}) bool {
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

	node, _ := list.getNodeByIndex(index)
	beforeNode := node.prev

	node.prev = newNode
	newNode.next = node
	newNode.prev = beforeNode
	if beforeNode != list.last {
		beforeNode.next = newNode
	} else {
		list.first = newNode
	}
	list.last.next = list.first
	list.first.prev = list.last
	list.size++
	return true
}

func (list *DoubleCircularLinkedList) Remove(index int) (interface{}, bool) {
	if !list.rangeCheck(index) {
		return nil, false
	}

	node, _ := list.getNodeByIndex(index)
	val := node.value

	if node == list.first {
		list.first = node.next
		if list.first != nil {
			list.first.prev = list.last
		}
		if list.last != nil {
			list.last.next = list.first
		}
	}
	if node == list.last {
		list.last = node.prev
		if list.first != nil {
			list.first.prev = list.last
		}
		if list.last != nil {
			list.last.next = list.first
		}
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	//
	node = nil
	list.size--
	return val, true
}

func (list *DoubleCircularLinkedList) Set(index int, v interface{}) bool {
	if node, ok := list.getNodeByIndex(index); ok {
		node.value = v
		return true
	}
	return false
}

func (list *DoubleCircularLinkedList) Get(index int) (interface{}, bool) {
	if node, ok := list.getNodeByIndex(index); ok {
		return node.value, true
	}
	return nil, false
}

func (list *DoubleCircularLinkedList) Contain(v interface{}) bool {
	return list.IndexOf(v) != -1
}

func (list *DoubleCircularLinkedList) IndexOf(v interface{}) int {
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

func (list *DoubleCircularLinkedList) Iterator() linearlist.Iterator {
	return NewIterator(list)
}

func (list *DoubleCircularLinkedList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}

func (list *DoubleCircularLinkedList) getNodeByIndex(index int) (*Node, bool) {
	if !list.rangeCheck(index) {
		return nil, false
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
	return node, true
}

func (list *DoubleCircularLinkedList) ValuesStartIndex(index int) []interface{} {
	if !list.rangeCheck(index) {
		return []interface{}{}
	}

	if list.size == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, list.size, list.size)
	node, _ := list.getNodeByIndex(index)
	for i := 0; i < list.size; i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}

func (list *DoubleCircularLinkedList) ValuesReverseStartIndex(index int) []interface{} {
	if !list.rangeCheck(index) {
		return []interface{}{}
	}

	if list.size == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, list.size, list.size)
	node, _ := list.getNodeByIndex(index)
	for i := 0; i < list.size; i++ {
		values[i] = node.value
		node = node.prev
	}
	return values
}
