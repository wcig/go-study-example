package arraystack

import (
	"go-app/algorithm/linearlist"
)

type ArrayStack struct {
	elements []interface{}
}

func New() *ArrayStack {
	return &ArrayStack{elements: []interface{}{}}
}

func (s *ArrayStack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

func (s *ArrayStack) Pop() (value interface{}, ok bool) {
	size := len(s.elements)
	if size == 0 {
		return nil, false
	}

	value = s.elements[size-1]
	s.elements = s.elements[:size-1]
	return value, true
}

// 获取栈顶元素但不弹出
func (s *ArrayStack) Peek() (value interface{}, ok bool) {
	if size := len(s.elements); size > 0 {
		return s.elements[size-1], true
	}
	return nil, false
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *ArrayStack) Size() int {
	return len(s.elements)
}

func (s *ArrayStack) Clear() {
	s.elements = []interface{}{}
}

func (s *ArrayStack) Values() []interface{} {
	values := make([]interface{}, len(s.elements), len(s.elements))
	size := s.Size()
	for i := 1; i <= size; i++ {
		values[size-i] = s.elements[i-1]
	}
	return values
}

func (s *ArrayStack) get(index int) (interface{}, bool) {
	if !s.rangeCheck(index) {
		return nil, false
	}
	return s.elements[index], true
}

func (s *ArrayStack) rangeCheck(index int) bool {
	return index >= 0 && index < len(s.elements)
}

func (s *ArrayStack) Iterator() linearlist.Iterator {
	return NewIterator(s)
}
