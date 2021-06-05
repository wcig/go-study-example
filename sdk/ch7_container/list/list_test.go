package list

import (
	"container/list"
	"fmt"
	"testing"
)

// container/list: 实现了双向链表

// 创建list函数:
// func New() *List

// list.List方法:
// func (l *List) Back() *Element  // 返回列表l的最后一个元素,如果l为空则返回nil
// func (l *List) Front() *Element // 返回列表l的第一个元素,如果l为空则返回nil
// func (l *List) Init() *List     // 初始化或列表l清空
// func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 在mark元素后插入元素v,如果列表l没有元素mark,则l不变 (mark不能为nil)
// func (l *List) InsertBefore(v interface{}, mark *Element) *Element // 在mark元素前插入元素v,如果列表l没有元素mark,则l不变 (mark不能为nil)
// func (l *List) Len() int // 返回列表l的元素数 (复杂度O(1))
// func (l *List) MoveBefore(e, mark *Element) // 移动元素e到mark元素之前,如果e后mark不是l的元素则l不变 (e和mark不能为nil)
// func (l *List) MoveToBack(e *Element) // 将元素e移动到列表l最后,e不属于l元素则l不变 (e不能为nil)
// func (l *List) MoveToFront(e *Element) // 将元素e移动到列表l最前,e不属于l元素则l不变 (e不能为nil)
// func (l *List) PushBack(v interface{}) *Element // 将v插入到列表l最后并返回该元素
// func (l *List) PushBackList(other *List) // 将列表里other插入到列表l最后 (l和other不能为nil)
// func (l *List) PushFront(v interface{}) *Element // 将v插入到列表l最前并返回该元素
// func (l *List) PushFrontList(other *List) // 将列表里other插入到列表l最前 (l和other不能为nil)
// func (l *List) Remove(e *Element) interface{} // 如果元素e是列表l的元素则移除,返回元素e的值 (e不能为nil)

// list.Element方法:
// func (e *Element) Next() *Element // 获取列表的下一个元素或nil
// func (e *Element) Prev() *Element // 获取列表的前一个元素或nil

func Test(t *testing.T) {
	l := list.New()

	l.Init()

	e4 := l.PushBack(4)
	printList(l)

	e1 := l.PushFront(1)
	printList(l)

	l.InsertBefore(3, e4)
	printList(l)

	l.InsertAfter(2, e1)
	printList(l)
	// output:
	// list: [4]
	// list: [1 4]
	// list: [1 3 4]
	// list: [1 2 3 4]
}

func printList(l *list.List) {
	var val []interface{}
	for e := l.Front(); e != nil; e = e.Next() {
		val = append(val, e.Value)
	}
	fmt.Println("list:", val)
}
