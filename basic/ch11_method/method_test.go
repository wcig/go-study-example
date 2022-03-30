package ch11_method

import (
	"fmt"
	"testing"
)

type user struct {
	Id   int
	Name string
}

func (u *user) print1() {
	fmt.Println("1-user:", u)
}

func (u user) print2() {
	fmt.Println("2-user:", u)
}

// method定义
func TestMethodDeclare(t *testing.T) {
	u1 := user{Id: 10, Name: "tom"}
	u1.print1() // 等价于: (&u1).print1(), go语言做了自动处理, 不需要像c语言那样处理
	u1.print2()

	u2 := &user{Id: 20, Name: "jerry"}
	u2.print1()
	u2.print2() // 等价于: (*u2).print2(), go语言做了自动处理, 不需要像c语言那样处理
}

// output:
// 1-user: &{10 tom}
// 2-user: {10 tom}
// 1-user: &{20 jerry}
// 2-user: {20 jerry}

func (u *user) modify1() {
	u.Id += 100
}

func (u user) modify2() {
	u.Id += 100
}

// receiver为值类型和引用类型区别
func TestMethodReceiver(t *testing.T) {
	u11 := user{Id: 1, Name: "Tom"}
	u11.modify1()
	fmt.Println("u11:", u11)

	u12 := user{Id: 1, Name: "Tom"}
	u11.modify2()
	fmt.Println("u12:", u12)

	u21 := &user{Id: 20, Name: "jerry"}
	u21.modify1()
	fmt.Println("u21:", u21)

	u22 := &user{Id: 20, Name: "jerry"}
	u22.modify2()
	fmt.Println("u22:", u22)
}

// output:
// u11: {101 Tom}
// u12: {1 Tom}
// u21: &{120 jerry}
// u22: &{20 jerry}

type num int

func (n *num) print() {
	fmt.Println("num:", *n)
}

// 类型别名
func TestTypeAlias(t *testing.T) {
	n := num(1)
	n.print() // num: 1
}

type aUser struct {
	name  string
	email string
}

func (u *aUser) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	aUser
	level string
}

// 嵌入类型method
func TestNestedStructMethod(t *testing.T) {
	ad := admin{
		aUser: aUser{
			name:  "Tom",
			email: "xx@gmail.com",
		},
		level: "super",
	}
	ad.aUser.notify() // 内部类型方法调用
	ad.notify()       // 外部类型方法调用
}

// output:
// Sending user email to Tom<xx@gmail.com>
// Sending user email to Tom<xx@gmail.com>

// func (a *admin) notify() {
// 	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
// }

// output:
// Sending user email to Tom<xx@gmail.com>
// Sending admin email to Tom<xx@gmail.com>

// 接口区别
type printer1 interface {
	print1()
}

type printer2 interface {
	print2()
}

func TestMethodInterface(t *testing.T) {
	var p11 printer1 = &user{1, "tom"}
	p11.print1()
	// var p12 printer1 = user{1, "tom"} // 错误
	// p12.print1()

	var p21 printer2 = user{2, "jerry"}
	p21.print2()
	var p22 printer2 = &user{2, "jerry"} // 允许
	p22.print2()

	// Output:
	// 1-user: &{1 tom}
	// 2-user: {2 jerry}
	// 2-user: {2 jerry}
}
