package ring

import (
	"container/ring"
	"fmt"
	"testing"
)

// container/ring
// 环是循环列表或环的元素。戒指没有开始或结束；指向任何环元素的指针用作对整个环的引用。空环表示为 nil 环指针。 Ring 的零值是一个值为 nil 的单元素环。

// 创建ring函数:
// func New(n int) *Ring

// ring.Ring方法:
// func (r *Ring) Do(f func(interface{})) // 对环r内每个函数执行函数f
// func (r *Ring) Len() int // 计算环r的元素个数 (执行效率与元素个数成比例)
// func (r *Ring) Link(s *Ring) *Ring // 将环r与环s连接起来,使得r.Next()等于s,返回r.Next()的原始值 (r不为空)
// func (r *Ring) Move(n int) *Ring // 在环r中向后(n<0)或向前(n>=0)移除n%r.Len()个元素并返回该环 (r不为空)
// func (r *Ring) Next() *Ring // 返回下一个环元素 (r不为空)
// func (r *Ring) Prev() *Ring // 返回前一个环元素 (r不为空)
// func (r *Ring) Unlink(n int) *Ring // 开始于r.Next(),从环r移除n%r.Len()个元素.如果n%r.Len()==0则r保持不变.返回结果为移除的子环 (r不为空)

func TestNew(t *testing.T) {
	r := ring.New(3)
	n := r.Len()
	fmt.Println("new ring len:", n)
}

func TestRingDo(t *testing.T) {
	r := ring.New(3)

	// 初始化ring的值
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	// 遍历ring每个元素并执行函数
	r.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})
	// output:
	// 0
	// 1
	// 2
}

func TestRingLink(t *testing.T) {
	r := ring.New(2)
	s := ring.New(3)

	lr := r.Len()
	ls := s.Len()

	for i := 0; i < lr; i++ {
		r.Value = i
		r = r.Next()
	}

	for j := 0; j < ls; j++ {
		s.Value = 10 + j
		s = s.Next()
	}

	rs := r.Link(s)
	rs.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})
	// output:
	// 1
	// 0
	// 10
	// 11
	// 12
}

func TestRingMove(t *testing.T) {
	r := ring.New(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	s := r.Move(3)

	fmt.Println("r: ")
	r.Do(func(i interface{}) {
		fmt.Printf("%d ", i.(int))
	})
	fmt.Println("\ns: ")
	s.Do(func(i interface{}) {
		fmt.Printf("%d ", i.(int))
	})
	// output:
	// r:
	// 0 1 2 3 4
	// s:
	// 3 4 0 1 2
}

func TestRingUnlink(t *testing.T) {
	r := ring.New(6)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int)) // 0 1 2 3 4 5
	})
	fmt.Println()

	r.Unlink(3)
	r.Do(func(p interface{}) {
		fmt.Printf("%d ", p.(int)) // 0 4 5
	})
}

func TestNextPrev(t *testing.T) {
	r := ring.New(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", r.Value.(int)) // 0 1 2 3 4
		r = r.Next()
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		r = r.Prev()
		fmt.Printf("%d ", r.Value.(int)) // 4 3 2 1 0
	}
}
