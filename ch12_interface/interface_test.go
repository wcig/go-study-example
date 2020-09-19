package ch12_interface

import (
	"fmt"
	"testing"
)

// 值接收者/指针接收者实现接口
func TestReceiverInterface(t *testing.T) {
	au := aUser{
		Name:  "tom",
		Email: "xxx@gmail.com",
	}
	au.notify()
	sendNotification(au)
	sendNotification(&au)

	bu := bUser{
		Name:  "jerry",
		Email: "yyy@gmail.com",
	}
	bu.notify()
	// sendNotification(bu) //  cannot use bu (type bUser) as type notifier in argument to sendNotification:
	sendNotification(&bu)
}

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type aUser struct {
	Name  string
	Email string
}

func (u aUser) notify() {
	fmt.Printf("send email to name:%s, email:%s\n", u.Name, u.Email)
}

type bUser struct {
	Name  string
	Email string
}

func (u *bUser) notify() {
	fmt.Printf("send email to name:%s, email:%s\n", u.Name, u.Email)
}

// 嵌入接口
type notifier2 interface {
	notify2()
	who
}

type who interface {
	name()
}

// 空接口
type object interface{}

func TestEmptyInterface(t *testing.T) {
	type user struct {
		Id   int
		Name string
	}

	var o object
	o = user{Id: 1, Name: "tom"}
	fmt.Printf("%T, %+v\n", o, o) // ch12_interface.user, {Id:1 Name:tom}
}

// 接口赋值是拷贝
type notifier3 interface {
	notify3()
	printer
}

type printer interface {
	print()
}

type cUser struct {
	Name  string
	Email string
}

func (u cUser) notify3() {
	fmt.Printf("send email to name:%s, email:%s\n", u.Name, u.Email)
}

func (u cUser) print() {
	fmt.Printf("name:%s\n", u.Name)
}

func TestInterfaceAssignment(t *testing.T) {
	u := cUser{
		Name:  "tom",
		Email: "xxx@gmail.com",
	}

	var p printer
	p = u
	fmt.Println("user:", u)    // user: {tom xxx@gmail.com}
	fmt.Println("printer:", p) // printer: {tom xxx@gmail.com}

	u.Name = "jerry"
	fmt.Println("user:", u)    // user: {jerry xxx@gmail.com}
	fmt.Println("printer:", p) // printer: {tom xxx@gmail.com}
}

// 接口nil
func TestInterfaceNil(t *testing.T) {
	var a interface{}
	fmt.Println(a == nil) // true

	var n *int = nil
	a = n
	fmt.Println(a == nil) // false
}
