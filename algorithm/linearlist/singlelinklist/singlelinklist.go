package singlelinklist

// 单链表
type SingleLinkList struct {
	first *Node
	size  int
}

type Node struct {
	data interface{}
	next *Node
}

func New() *SingleLinkList {
	return &SingleLinkList{
		first: nil,
		size:  0,
	}
}

func (l *SingleLinkList) Size() int {
	return l.size
}

func (l *SingleLinkList) IsEmpty() bool {
	return l.Size() == 0
}

func (l *SingleLinkList) Clear() {
	l.first = nil
	l.size = 0
}

func (l *SingleLinkList) Values() []interface{} {
	if l.IsEmpty() {
		return []interface{}{}
	}

	newData := make([]interface{}, 0, l.size)
	node := l.first
	for node.next != nil {
		newData = append(newData, node.data)
		node = node.next
	}
	return newData
}

func (l *SingleLinkList) Add(v interface{}) {
	addNode := &Node{
		data: v,
		next: nil,
	}

	l.size++
	if l.first == nil {
		l.first = addNode
	} else {
		node := l.first
		for node.next != nil {
			node = node.next
		}
		node.next = addNode
	}
}
