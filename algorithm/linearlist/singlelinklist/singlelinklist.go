package singlelinklist

import "go-app/algorithm/linearlist"

// 单链表
type SingleLinkList struct {
	first *Node
	last  *Node
	size  int
}

type Node struct {
	data interface{}
	next *Node
}

func New() *SingleLinkList {
	return &SingleLinkList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (list *SingleLinkList) Size() int {
	return list.size
}

func (list *SingleLinkList) IsEmpty() bool {
	return list.Size() == 0
}

func (list *SingleLinkList) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *SingleLinkList) Values() []interface{} {
	if list.IsEmpty() {
		return []interface{}{}
	}

	values := make([]interface{}, 0, list.size)
	node := list.first
	for {
		values = append(values, node.data)
		node = node.next
		if node == nil {
			break
		}
	}
	return values
}

func (list *SingleLinkList) Add(v interface{}) {
	newNode := &Node{
		data: v,
		next: nil,
	}

	if list.IsEmpty() {
		list.first = newNode
		list.last = newNode
		list.size++
		return
	}
	beforeNode := list.first
	for {
		if beforeNode.next != nil {
			beforeNode = beforeNode.next
		} else {
			break
		}
	}
	beforeNode.next = newNode
	list.last = newNode
	list.size++
}

func (list *SingleLinkList) Insert(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		if index == list.size {
			list.Add(v)
			return true
		}
		return false
	}

	newNode := &Node{
		data: v,
		next: nil,
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

func (list *SingleLinkList) Remove(index int) (interface{}, bool) {
	// TODO implement me
	panic("implement me")
}

func (list *SingleLinkList) Set(index int, v interface{}) bool {
	if !list.rangeCheck(index) {
		return false
	}

	newNode := &Node{
		data: v,
		next: nil,
	}

	beforeNode := list.first
	for i := 0; i < index-1; i++ {
		beforeNode = beforeNode.next
	}
	node := beforeNode.next
	newNode.next = node.next
	beforeNode.next = newNode
	return true
}

func (list *SingleLinkList) Get(index int) (interface{}, bool) {
	// TODO implement me
	panic("implement me")
}

func (list *SingleLinkList) Contain(v interface{}) bool {
	if list.IsEmpty() {
		return false
	}

	node := list.first
	if node.data == v {
		return true
	}
	for {
		if node.next != nil {
			node = node.next
			if node.data == v {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func (list *SingleLinkList) IndexOf(v interface{}) int {
	// TODO implement me
	panic("implement me")
}

func (list *SingleLinkList) Iterator() *linearlist.Iterator {
	// TODO implement me
	panic("implement me")
}

func (list *SingleLinkList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}
