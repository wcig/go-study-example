package doublelinklist

import (
	"go-app/algorithm/linearlist"
)

// 双向链表
// 单链表
type DoubleLinkList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func New() *DoubleLinkList {
	return &DoubleLinkList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (list *DoubleLinkList) Size() int {
	return list.size
}

func (list *DoubleLinkList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoubleLinkList) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *DoubleLinkList) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, node := 0, list.first; i < list.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *DoubleLinkList) Add(v interface{}) {
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
func (list *DoubleLinkList) Insert(index int, v interface{}) bool {
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

	if index == 0 {
		node := list.first
		newNode.next = node
		node.prev = newNode
		list.first = newNode
	} else {
		node := list.getNodeByIndex(index)
		beforeNode := node.prev
		beforeNode.next = newNode
		newNode.prev = beforeNode
		newNode.next = node
		node.prev = newNode
	}
	list.size++
	return true
}

func (list *DoubleLinkList) Remove(index int) (interface{}, bool) {
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

func (list *DoubleLinkList) Set(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		return false
	}

	node := list.getNodeByIndex(index)
	node.value = v
	return true
}

func (list *DoubleLinkList) Get(index int) (interface{}, bool) {
	if !list.rangeCheck(index) {
		return nil, false
	}

	node := list.getNodeByIndex(index)
	return node.value, true
}

func (list *DoubleLinkList) Contain(v interface{}) bool {
	return list.IndexOf(v) != -1
}

func (list *DoubleLinkList) IndexOf(v interface{}) int {
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

func (list *DoubleLinkList) Iterator() linearlist.Iterator {
	return NewIterator(list)
}

func (list *DoubleLinkList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}

func (list *DoubleLinkList) getNodeByIndex(index int) *Node {
	if !list.rangeCheck(index) {
		return nil
	}

	var node *Node
	if list.size-index < index {
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
