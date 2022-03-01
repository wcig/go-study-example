package singlelinklist

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
	for node != nil {
		values = append(values, node.data)
		node = node.next
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

// index从0开始, 当index为列表size时,直接添加到末尾 (index<0 || index>size返回false)
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

// index从0开始 (index<0 || index>=size返回false)
func (list *SingleLinkList) Remove(index int) (interface{}, bool) {
	if list.IsEmpty() {
		return nil, false
	}

	if !list.rangeCheck(index) {
		return nil, false
	}

	// 删除第一个元素
	if index == 0 {
		node := list.first
		val := node.data
		list.first = node.next
		if list.first == nil {
			list.last = nil
		}
		list.size--
		return val, true
	}

	var val interface{}
	beforeNode := list.first
	for i := 0; i < index-1; i++ {
		beforeNode = beforeNode.next
	}
	node := beforeNode.next
	val = node.data
	if node == list.last {
		// 删除最后一个元素
		beforeNode.next = nil
		list.last = beforeNode
	} else {
		// 删除中间元素
		beforeNode.next = node.next
	}
	list.size--
	return val, true
}

// index从0开始 (index<0 || index>=size返回false)
func (list *SingleLinkList) Set(index int, v interface{}) bool {
	if list.IsEmpty() {
		return false
	}

	if !list.rangeCheck(index) {
		return false
	}

	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.data = v
	return true
}

// index从0开始  (index<0 || index>=size返回nil,false)
func (list *SingleLinkList) Get(index int) (interface{}, bool) {
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
	return node.data, true
}

// 返回list是否存在元素与v相等
func (list *SingleLinkList) Contain(v interface{}) bool {
	return list.IndexOf(v) != -1
}

// 返回list元素红第一次出现v的索引index, index从0开始 (存在返回index,不存在返回-1)
func (list *SingleLinkList) IndexOf(v interface{}) int {
	if list.IsEmpty() {
		return -1
	}

	node := list.first
	if node.data == v {
		return 0
	}
	for i := 0; i < list.size-1; i++ {
		node = node.next
		if node.data == v {
			return i + 1
		}
	}
	return -1
}

func (list *SingleLinkList) Iterator() *Iterator {
	return NewIterator(list)
}

func (list *SingleLinkList) rangeCheck(index int) bool {
	return index >= 0 && index < list.size
}
