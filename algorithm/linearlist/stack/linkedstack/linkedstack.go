package linkedstack

import (
	"go-app/algorithm/linearlist"
)

// 链式堆栈
type LinkedStack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func New() *LinkedStack {
	return &LinkedStack{
		top:  nil,
		size: 0,
	}
}

// 从栈顶推入元素 (底层从链表表头插入元素)
func (s *LinkedStack) Push(value interface{}) {
	newElement := &Element{
		value: value,
		next:  s.top,
	}
	s.top = newElement
	s.size++
}

// 从栈顶弹出元素 (底层从链表表尾弹出元素)
func (s *LinkedStack) Pop() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}

	element := s.top
	value = element.value
	s.top = s.top.next
	s.size--
	element = nil
	return value, true
}

// 获取栈顶元素但不弹出
func (s *LinkedStack) Peek() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}
	return s.top.value, true
}

func (s *LinkedStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedStack) Size() int {
	return s.size
}

func (s *LinkedStack) Clear() {
	s.top = nil
	s.size = 0
}

func (s *LinkedStack) Values() []interface{} {
	values := make([]interface{}, s.size, s.size)
	for i, node := 0, s.top; i < s.size; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (s *LinkedStack) Iterator() linearlist.Iterator {
	return NewIterator(s)
}

func (s *LinkedStack) get(index int) (interface{}, bool) {
	if !s.rangeCheck(index) {
		return nil, false
	}

	e := s.top
	for i := 0; i < index; i++ {
		e = e.next
	}
	return e.value, true
}

func (s *LinkedStack) rangeCheck(index int) bool {
	return index >= 0 && index < s.size
}
