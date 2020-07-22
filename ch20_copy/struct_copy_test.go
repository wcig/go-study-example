package ch20_copy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
)

// 浅拷贝与深拷贝 (shallow copy, deep copy)
type user1 struct {
	id   int
	name string
}

type user2 struct {
	id     int
	name   string
	phones []string
}

type Cat struct {
	Age     int
	Name    string
	Friends []string
}

// 成员变量为基础数据类型的结构体: 浅拷贝与深拷贝一致
func TestWithBasicDataType(t *testing.T) {
	u1 := user1{1, "tom"}

	// 浅拷贝
	u2 := u1
	fmt.Println(u1 == u2) // true

	u2.name = "jerry"
	fmt.Println(u1 == u2) // false
	fmt.Println(u1, u2)   // {1 tom} {1 jerry}

	// 深拷贝
	u3 := user1{u1.id, u1.name}
	fmt.Println(u1 == u3) // true

	u3.name = "jerry"
	fmt.Println(u1 == u3) // false
	fmt.Println(u1, u3)   // {1 tom} {1 jerry}
}

// 成员变量包含指针类型的结构体: 浅拷贝与深拷贝不一致
func TestWithPointerType(t *testing.T) {
	u1 := user2{1, "tom", []string{"001", "002", "003"}}

	// 浅拷贝
	u2 := u1
	// fmt.Println(u1 == u2) // 错误: invalid operation: u1 == u2 (struct containing []string cannot be compared)

	u2.phones[0] = "100" // {1 tom [100 002 003]} {1 tom [100 002 003]}
	fmt.Println(u1, u2)

	// 深拷贝
	phones := make([]string, len(u1.phones))
	copy(phones, u1.phones)
	u3 := user2{u1.id, u1.name, phones}

	u3.phones[0] = "200"
	fmt.Println(u1, u3) // {1 tom [100 002 003]} {1 tom [200 002 003]}

	// 自定义函数深拷贝
	u4 := copyUser2(u1)
	u4.phones[0] = "300"
	fmt.Println(u1, u4) // {1 tom [100 002 003]} {1 tom [300 002 003]}
}

func copyUser2(u user2) user2 {
	phones := make([]string, len(u.phones))
	copy(phones, u.phones)
	return user2{u.id, u.name, phones}
}

// 使用库实现深拷贝 (struct成员变量为非导出类型,使用reflect的库不生效. 测试成员变量为导出类型时这里依然不生效)
func TestDeepCopyWithLibrary1(t *testing.T) {
	u1 := user2{1, "tom", []string{"001", "002", "003"}}
	u2 := user2{}
	err := copier.Copy(&u2, &u1)
	assert.True(t, err == nil)

	fmt.Println(u2)
	u2.phones = append(u2.phones, "100") // {1 tom [001 002 003]} {0  [100]}
	fmt.Println(u1, u2)
}

// "github.com/jinzhu/copier"库不建议使用
func TestDeepCopyWithLibrary2(t *testing.T) {
	wilson := Cat{7, "Wilson", []string{"Tom", "Tabata", "Willie"}}
	nikita := Cat{}
	copier.Copy(&nikita, &wilson)

	nikita.Friends[0] = "Syd"
	fmt.Println(wilson, nikita) // {7 Wilson [Syd Tabata Willie]} {7 Wilson [Syd Tabata Willie]}
}

// 使用json方式
func TestWithJson(t *testing.T) {
	// 结构体变量为非导出类型
	u1 := user2{1, "tom", []string{"001", "002", "003"}}
	u2 := user2{}

	bytes1, err := json.Marshal(&u1)
	assert.True(t, err == nil)

	fmt.Println(string(bytes1)) // {}

	err = json.Unmarshal(bytes1, &u2)
	assert.True(t, err == nil)

	fmt.Println(u1, u2) // {1 tom [001 002 003]} {0  []}

	// 结构体变量为非导出类型
	type user struct {
		Id     int
		Name   string
		Phones []string
	}

	u3 := user{1, "tom", []string{"001", "002", "003"}}
	u4 := user{}

	bytes3, err := json.Marshal(&u3)
	assert.True(t, err == nil)

	err = json.Unmarshal(bytes3, &u4)
	assert.True(t, err == nil)

	fmt.Println(u3, u4)                         // {1 tom [001 002 003]} {1 tom [001 002 003]}
	fmt.Println(len(u3.Phones), len(u4.Phones)) // 3 3

	u4.Phones = append(u4.Phones, "004")
	fmt.Println(u3, u4)                         // {1 tom [001 002 003]} {1 tom [001 002 003 004]}
	fmt.Println(len(u3.Phones), len(u4.Phones)) // 3 4
}

// gob方式 (struct的成员变量必须为导出类型)
func TestWithGob(t *testing.T) {
	type user struct {
		Id     int
		Name   string
		Phones []string
	}

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	dec := gob.NewDecoder(&buffer)

	u1 := user{1, "tom", []string{"001", "002", "003"}}
	err := enc.Encode(&u1)
	assert.True(t, err == nil)

	u2 := user{}
	err = dec.Decode(&u2)
	assert.True(t, err == nil)

	fmt.Println(u1, u2)                         // {1 tom [001 002 003]} {1 tom [001 002 003]}
	fmt.Println(len(u1.Phones), len(u2.Phones)) // 3 3

	u2.Phones = append(u2.Phones, "004")
	fmt.Println(u1, u2)                         // {1 tom [001 002 003]} {1 tom [001 002 003 004]}
	fmt.Println(len(u1.Phones), len(u2.Phones)) // 3 4
}
