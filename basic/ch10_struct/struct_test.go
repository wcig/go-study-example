package ch10_struct

import (
	"fmt"
	"testing"
)

// struct定义
type user struct {
	Id   int
	Name string
}

// struct初始化
func TestStructInit(t *testing.T) {
	// 初始化结构体值
	// 1.先声明再赋值
	var u1 user
	u1.Id = 10
	u1.Name = "tom"
	fmt.Printf("%+v\n", u1) // {Id:10 Name:tom}

	// 2.声明同时赋值
	u2 := user{Id: 20, Name: "jerry"}
	fmt.Printf("%+v\n", u2) // {Id:20 Name:jerry}

	// 初始化结构体指针
	u3 := &user{Id: 30, Name: "jack"}
	fmt.Printf("%+v\n", u3) // &{Id:30 Name:jack}
}

// struct声明初始值
func TestStructDeclare(t *testing.T) {
	var u1 user
	fmt.Printf("%+v\n", u1) // {Id:0 Name:}

	var u2 *user
	fmt.Printf("%+v\n", u2) // <nil>

	var u3 = &user{}
	fmt.Printf("%+v\n", u3) // &{Id:0 Name:}
}

// struct成员变量
func TestStructField(t *testing.T) {
	// 结构体对象访问成员变量
	u1 := user{Id: 10, Name: "tom"}
	fmt.Printf("id:%d, name:%s\n", u1.Id, u1.Name) // id:10, name:tom
	u1.Id = 100

	// 结构体指针访问成员变量
	u2 := &user{Id: 20, Name: "jerry"}
	fmt.Printf("id:%d, name:%s\n", u2.Id, u2.Name) // id:20, name:jerry
	u2.Id = 200

	// 此时声明为结构体指针，需初始化
	var u3 *user
	fmt.Printf("id:%d, name:%s\n", u3.Id, u3.Name) // panic: runtime error: invalid memory address or nil pointer dereference
}

// struct作为参数
func TestStructAsParam(t *testing.T) {
	u1 := user{Id: 10, Name: "tom"}
	modifyWithStructObject(u1)
	fmt.Printf("after modify: %+v\n", u1)

	u2 := &user{Id: 20, Name: "jerry"}
	modifyWithStructPointer(u2)
	fmt.Printf("%+v\n", u2)
}

// output:
// modify: {Id:100 Name:tom}
// after modify: {Id:10 Name:tom}
// modify: &{Id:200 Name:jerry}
// &{Id:200 Name:jerry}

func modifyWithStructObject(u user) {
	u.Id = 100
	fmt.Printf("modify: %+v\n", u)
}

func modifyWithStructPointer(u *user) {
	u.Id = 200
	fmt.Printf("modify: %+v\n", u)
}

// struct比较
func TestStructCompare(t *testing.T) {
	u1 := user{Id: 10, Name: "tom"}
	u2 := user{Id: 10, Name: "tom"}
	fmt.Println(u1 == u2) // true

	u3 := &user{Id: 20, Name: "jerry"}
	u4 := &user{Id: 20, Name: "jerry"}
	u5 := u3
	fmt.Println(u3 == u4) // false
	fmt.Println(u5 == u3) // true

	type xUser struct {
		Id   int
		Name string
	}
	// var x1 xUser
	// fmt.Println(x1 == u1) // 不同类型不能比较

	type person struct {
		user
		phone []string
	}
	// var p1 person
	// var p2 person
	// fmt.Println(p1 == p2) // invalid operation: p1 == p2 (struct containing []string cannot be compared)
}

// 匿名结构体
func TestAnonymousStruct(t *testing.T) {
	u1 := struct {
		Id   int
		Name string
	}{
		Id:   10,
		Name: "tom",
	}
	fmt.Printf("%+v\n", u1) // {Id:10 Name:tom}

	type xUser struct {
		int
		string
	}
	x1 := xUser{
		int:    20,
		string: "jerry",
	}
	fmt.Printf("%+v\n", x1) // {int:20 string:jerry}
}

// 嵌套结构体
func TestNestedStruct(t *testing.T) {
	type person1 struct {
		user
		Phone string
	}

	p11 := person1{
		user: user{
			Id:   10,
			Name: "tom",
		},
		Phone: "001",
	}
	fmt.Printf("%+v\n", p11) // {user:{Id:10 Name:tom} Phone:001}

	p12 := person1{
		Phone: "002",
	}
	u12 := user{Id: 20, Name: "jerry"}
	p12.user = u12
	fmt.Printf("%+v\n", p12) // {user:{Id:20 Name:jerry} Phone:002}

	type person2 struct {
		Id      int
		Name    string
		Contact struct {
			Phone string
			City  string
		}
	}

	p21 := person2{
		Id:   30,
		Name: "A",
		Contact: struct {
			Phone string
			City  string
		}{
			Phone: "003",
			City:  "AAA",
		},
	}
	fmt.Printf("%+v\n", p21) // {Id:30 Name:A Contact:{Phone:003 City:AAA}}

	p22 := person2{
		Id:   40,
		Name: "B",
	}
	p22.Contact.Phone = "004"
	p22.Contact.City = "BBB"
	fmt.Printf("%+v\n", p22) // {Id:40 Name:B Contact:{Phone:004 City:BBB}}
}

// 特殊嵌套结构体: 外层与内层存在同名字段
func TestSpecialStruct(t *testing.T) {
	a := A1{
		Name: "A",
		B1: B1{
			Name: "B",
		},
	}
	fmt.Println(a)                 // {{B} A}
	fmt.Println(a.Name, a.B1.Name) // A B

	a2 := A2{
		B2{
			Name: "B",
		},
	}
	fmt.Println(a2)                  // {{B}}
	fmt.Println(a2.Name, a2.B2.Name) // B B

	a3 := A3{
		B3{
			Name: "B",
		},
		C3{
			Name: "C",
		},
	}
	fmt.Println(a3) // {{B} {C}}
	// fmt.Println(a3.Name) // 错误
	fmt.Println(a3.B3.Name, a3.C3.Name) // B C
}

type A1 struct {
	B1
	Name string
}
type B1 struct {
	Name string
}

type A2 struct {
	B2
}
type B2 struct {
	Name string
}

type A3 struct {
	B3
	C3
}
type B3 struct {
	Name string
}
type C3 struct {
	Name string
}

// struct强转
func TestStructConvert(t *testing.T) {
	type user1 struct {
		UserId   int    `json:"user_id"`
		UserName string `json:"user_name"`
	}
	type user2 struct {
		UserId   int    `json:"userId"`
		UserName string `json:"userName"`
	}

	u1 := user1{UserId: 10, UserName: "tom"}
	u2 := user2(u1)
	fmt.Printf("%+v\n", u1) // {UserId:10 UserName:tom}
	fmt.Printf("%+v\n", u2) // {UserId:10 UserName:tom}

	u3 := &user1{UserId: 20, UserName: "jerry"}
	u4 := user2(*u3)
	u5 := &u4
	fmt.Printf("%+v\n", u3) // &{UserId:20 UserName:jerry}
	fmt.Printf("%+v\n", u4) // {UserId:20 UserName:jerry}
	fmt.Printf("%+v\n", u5) // &{UserId:20 UserName:jerry}
}
