package ch8_map

import (
	"fmt"
	"testing"
)

// map初始化
func TestMapInit(t *testing.T) {
	// 声明 (此时map为nil不能赋值)
	var m1 map[int]string
	// m1[1] = "one" // panic: assignment to entry in nil map
	fmt.Println("m1=", m1, ", len=", len(m1))

	// 声明并初始化
	m2 := map[int]string{}
	m2[2] = "two" // 等价于 m2 := map[int]string{2: "two"}
	fmt.Println("m2=", m2, ", len=", len(m2))

	// make初始化
	var m3 = make(map[int]string) // 等价于 m3 := make(map[int]string)
	m3[3] = "three"
	fmt.Println("m3=", m3, ", len=", len(m3))

	// make初始化并设置容量
	m4 := make(map[int]string, 3)
	m4[4] = "four"
	fmt.Println("m4=", m4, ", len=", len(m4))
}

// output:
// m1= map[] , len= 0
// m2= map[2:two] , len= 1
// m3= map[3:three] , len= 1
// m4= map[4:four] , len= 1

// map添加元素
func TestMapAppend(t *testing.T) {
	m := make(map[int]string)
	m[1] = "one"
	fmt.Println("m=", m, ", len=", len(m))

	m[1] = "1"
	fmt.Println("m=", m, ", len=", len(m))

	m[2] = "two"
	fmt.Println("m=", m, ", len=", len(m))
}

// map删除元素
func TestMapDelete(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Println("m=", m, ", len=", len(m))

	delete(m, 1)
	fmt.Println("m=", m, ", len=", len(m))

	// 删除不存在的元素不会报错
	delete(m, 1)
	fmt.Println("m=", m, ", len=", len(m))
}

// map遍历元素
func TestMapRange1(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	for key, val := range m {
		fmt.Println("key=", key, ", val=", val)
	}
}

// map遍历元素时修改元素对应的值
func TestMapRange2(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// 此时val为map真实val的拷贝，所以对于基础数据类型的val不能修改成功，对于指针类型可以
	for _, val := range m {
		val += "+"
	}
	fmt.Println("m=", m) // m= map[1:one 2:two 3:three]

	// 通过key修改值
	for key := range m {
		m[key] += "-"
	}
	fmt.Println("m=", m) // m= map[1:one- 2:two- 3:three-]
}

// map: 判断key是否存在
func TestMapKeyExists(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// key存在则返回对应值+true
	val1, exists := m[1]
	fmt.Println(val1, exists) // one true

	// key不存在则返回零值+false
	val2, exists := m[0]
	fmt.Println(val2, exists) //  false
}

// map是引用类型，作为参数传递可以被修改
func TestMapAsParam(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	modify(m)
	fmt.Println("m=", m) // m= map[1:one- 2:two- 3:three-]
}

func modify(m map[int]string) {
	for key := range m {
		m[key] += "-"
	}
}

// map作为返回值
func TestMapAsReturn(t *testing.T) {
	m1 := create1()
	fmt.Println("m1=", m1, ", len=", len(m1))

	// 注意: 当map作为返回值如果返回nil, 此时map不能添加元素
	m2 := create2()
	// m2[2] = "two" // panic: assignment to entry in nil map
	fmt.Println("m2=", m2, ", len=", len(m2))

	m3 := create3()
	m3[3] = "three"
	fmt.Println("m3=", m3, ", len=", len(m3))
}

// output:
// m1= map[1:one 2:two 3:three] , len= 3
// m2= map[] , len= 0
// m3= map[3:three] , len= 1

func create1() map[int]string {
	return map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
}

func create2() map[int]string {
	return nil
}

func create3() map[int]string {
	return map[int]string{}
}
