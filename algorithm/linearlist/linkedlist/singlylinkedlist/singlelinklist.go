package singlylinkedlist

import "go-app/algorithm/linearlist"

// 单向链表
type SinglyLinkedList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	value interface{}
	next  *Node
}

func New() *SinglyLinkedList {
	return &SinglyLinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (list *SinglyLinkedList) Size() int {
	return list.size
}

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *SinglyLinkedList) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *SinglyLinkedList) Values() []interface{} {
	if list.IsEmpty() {
		return []interface{}{}
	}

	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *SinglyLinkedList) Add(v interface{}) {
	newNode := &Node{
		value: v,
		next:  nil,
	}

	if list.size == 0 {
		list.first = newNode
		list.last = newNode
	} else {
		list.last.next = newNode
		list.last = newNode
	}
	list.size++
}

// index从0开始, 当index为列表size时,直接添加到末尾 (index<0 || index>size返回false)
func (list *SinglyLinkedList) Insert(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		if index == list.size {
			list.Add(v)
			return true
		}
		return false
	}

	newNode := &Node{
		value: v,
		next:  nil,
	}

	if index == 0 {
		newNode.next = list.first
		list.first = newNode
	} else {
		beforeNode := list.first
		for i := 0; i < index-1; i++ {
			beforeNode = beforeNode.next
		}
		node := beforeNode.next
		beforeNode.next = newNode
		newNode.next = node
	}
	list.size++
	return true
}

// index从0开始 (index<0 || index>=size返回false)
func (list *SinglyLinkedList) Remove(index int) (interface{}, bool) {
	if list.size == 0 {
		return nil, false
	}

	if !list.rangeCheck(index) {
		return nil, false
	}

	var beforeNode *Node
	node := list.first
	for i := 0; i < index; i++ {
		beforeNode = node
		node = node.next
	}
	val := node.value

	if node == list.first {
		list.first = node.next
	}
	if node == list.last {
		list.last = beforeNode
	}
	if beforeNode != nil {
		beforeNode.next = node.next
	}
	node = nil
	list.size--
	return val, true
}

// index从0开始 (index<0 || index>=size返回false)
func (list *SinglyLinkedList) Set(index int, v interface{}) bool {
	if list.size == 0 {
		return false
	}

	if !list.rangeCheck(index) {
		return false
	}

	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.value = v
	return true
}

// index从0开始  (index<0 || index>=size返回nil,false)
func (list *SinglyLinkedList) Get(index int) (interface{}, bool) {
	if list.IsEmpty() {
		return nil, false
	}

	if !list.rangeCheck(index) {
		return nil, false
	}

	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value, true
}

// 返回list是否存在元素与v相等
func (list *SinglyLinkedList) Contain(v interface{}) bool {
	return list.IndexOf(v) != -1
}

// 返回list元素红第一次出现v的索引index, index从0开始 (存在返回index,不存在返回-1)
func (list *SinglyLinkedList) IndexOf(v interface{}) int {
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

func (list *SinglyLinkedList) Iterator() linearlist.Iterator {
	return NewIterator(list)
}

func (list *SinglyLinkedList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}
