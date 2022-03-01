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
		var beforeNode *Node
		node := list.first
		for i := 0; i < index; i++ {
			beforeNode = node.next
			node = node.next
		}
		beforeNode.next = newNode
		newNode.prev = beforeNode
		newNode.next = node
		node.prev = newNode
	}
	list.size++
	return true
}

func (list *DoubleLinkList) Remove(index int) (interface{}, bool) {
	// TODO implement me
	panic("implement me")
}

func (list *DoubleLinkList) Set(index int, v interface{}) bool {
	// TODO implement me
	panic("implement me")
}

func (list *DoubleLinkList) Get(index int) (interface{}, bool) {
	// TODO implement me
	panic("implement me")
}

func (list *DoubleLinkList) Contain(v interface{}) bool {
	// TODO implement me
	panic("implement me")
}

func (list *DoubleLinkList) IndexOf(v interface{}) int {
	// TODO implement me
	panic("implement me")
}

func (list *DoubleLinkList) Iterator() linearlist.Iterator {
	// TODO implement me
	panic("implement me")
}

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func (list *DoubleLinkList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}
