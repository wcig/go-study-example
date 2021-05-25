package ch13_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TypeOf
func TestTypeOf(t *testing.T) {
	u := User{Id: 1, Name: "tom"}

	_type := reflect.TypeOf(u)
	fmt.Println("user type:", _type)

	num := _type.NumField()
	fmt.Println("user filed num:", num)

	for i := 0; i < num; i++ {
		filed := _type.Field(i)
		fmt.Printf("user filed: %d, %+v\n", i, filed)
	}
}

// output:
// user type: ch13_reflect.User
// user filed num: 2
// user filed: 0, {Name:Id PkgPath: Type:int Tag:json:"id" Offset:0 Index:[0] Anonymous:false}
// user filed: 1, {Name:Name PkgPath: Type:string Tag:json:"name" Offset:8 Index:[1] Anonymous:false}

// ValueOf
func TestValueOf(t *testing.T) {
	u := User{Id: 1, Name: "tom"}

	_value := reflect.ValueOf(u)
	fmt.Println("user value:", _value)

	num := _value.NumField()
	fmt.Println("user value num:", num)

	for i := 0; i < num; i++ {
		value := _value.Field(i)
		fmt.Printf("user value: %d, %+v\n", i, value)
	}
}

// output:
// user value: {1 tom}
// user value num: 2
// user value: 0, 1
// user value: 1, tom

// 获取成员变量值
func TestGetFileValue(t *testing.T) {
	u := User{Id: 1, Name: "tom"}
	_value := reflect.ValueOf(u)

	// Filed
	idVal1 := _value.Field(0)
	id1 := idVal1.Interface().(int)

	nameVal1 := _value.Field(1)
	name1 := nameVal1.Interface().(string)
	fmt.Println("Filed:", reflect.TypeOf(id1).String(), id1, reflect.TypeOf(name1).String(), name1)

	// FieldByName
	idVal2 := _value.FieldByName("Id")
	id2 := idVal2.Interface().(int)

	nameVal2 := _value.FieldByName("Name")
	name2 := nameVal2.Interface().(string)
	fmt.Println("FieldByName:", reflect.TypeOf(id2).String(), id2, reflect.TypeOf(name2).String(), name2)

	// get field name
	_type := reflect.TypeOf(u)
	for i := 0; i < _type.NumField(); i++ {
		fmt.Println("filed:", i, _type.Field(i).Name)
	}
}

// output:
// Filed: int 1 string tom
// FieldByName: int 1 string tom
// filed: 0 Id
// filed: 1 Name

// 私有成员变量: 私有成员变量类型可以获取,但是值不能获取
func TestPrivateField(t *testing.T) {
	type user struct {
		Id   int
		name string
	}
	u := user{Id: 1, name: "tom"}

	_type := reflect.TypeOf(u)
	for i := 0; i < _type.NumField(); i++ {
		fmt.Println("filed:", i, _type.Field(i).Name)
	}

	_value := reflect.ValueOf(u)
	for i := 0; i < _value.NumField(); i++ {
		fmt.Println("value:", i, _value.Field(i).Interface())
	}

	iv := _value.FieldByName("Id")
	fmt.Println("id:", iv.Interface().(int))

	// nv := _value.FieldByName("name")
	// fmt.Println("name:", nv.Interface().(string)) // error
}

// output:
// filed: 0 Id
// filed: 1 name
// value: 0 1
// --- FAIL: TestPrivateField (0.00s)
// panic: reflect.Value.Interface: cannot return value obtained from unexported field or method [recovered]
//	panic: reflect.Value.Interface: cannot return value obtained from unexported field or method

// 通过反射获取方法
func (u User) Print1() {
	fmt.Printf("Print1 user id:%d, name:%s\n", u.Id, u.Name)
}

func (u User) print2() {
	fmt.Printf("print2 user id:%d, name:%s\n", u.Id, u.Name)
}

func (u User) Print3(prefix string) {
	fmt.Printf("Print3 %s user id:%d, name:%s\n", prefix, u.Id, u.Name)
}

func (u *User) Print4() {
	fmt.Printf("Print4 user id:%d, name:%s\n", u.Id, u.Name)
}

// 私有方法不能通过反射获取
func TestGetMethod1(t *testing.T) {
	u := User{Id: 1, Name: "tom"}

	_type := reflect.TypeOf(u)
	for i := 0; i < _type.NumMethod(); i++ {
		fmt.Println("method:", i, _type.Method(i))
	}

	p1, has := _type.MethodByName("Print1")
	fmt.Println(p1, has)
	p2, has := _type.MethodByName("print2")
	fmt.Println(p2, has)
}

// output:
// method: 0 {Print1  func(ch13_reflect.User) <func(ch13_reflect.User) Value> 0}
// method: 1 {Print3  func(ch13_reflect.User, string) <func(ch13_reflect.User, string) Value> 1}
// {Print1  func(ch13_reflect.User) <func(ch13_reflect.User) Value> 0} true
// {  <nil> <invalid Value> 0} false

type Object struct{}

func (o Object) Ok1() {
	fmt.Println("Ok1")
}

func (o *Object) Ok2() {
	fmt.Println("Ok2")
}

// 接收者为值类型方法可以通过反射只能获取值类型方法, 接收者类型为引用方法可以获取值类型和引用类型方法
func TestGetMethod2(t *testing.T) {
	var o Object

	_type1 := reflect.TypeOf(o)
	for i := 0; i < _type1.NumMethod(); i++ {
		fmt.Println("method1:", i, _type1.Method(i))
	}

	_type2 := reflect.TypeOf(&o)
	for i := 0; i < _type2.NumMethod(); i++ {
		fmt.Println("method2:", i, _type2.Method(i))
	}
}

// output:
// method1: 0 {Ok1  func(ch13_reflect.Object) <func(ch13_reflect.Object) Value> 0}
// method2: 0 {Ok1  func(*ch13_reflect.Object) <func(*ch13_reflect.Object) Value> 0}
// method2: 1 {Ok2  func(*ch13_reflect.Object) <func(*ch13_reflect.Object) Value> 1}

// 通过反射调用方法
func TestCallMethod1(t *testing.T) {
	u := User{Id: 1, Name: "tom"}
	_type := reflect.ValueOf(u)

	p1 := _type.MethodByName("Print1")
	if p1.IsValid() {
		p1.Call(nil) // Print1 user id:1, name:tom
	}

	p2 := _type.MethodByName("print2")
	fmt.Println("print2", p2.IsValid()) // print2 false

	p3 := _type.MethodByName("Print3")
	if p3.IsValid() {
		args := []reflect.Value{reflect.ValueOf("hello")}
		p3.Call(args) // Print3 hello user id:1, name:tom
	}
}

func TestCallMethod2(t *testing.T) {
	u := User{Id: 1, Name: "tom"}

	v1 := reflect.ValueOf(u)
	p11 := v1.MethodByName("Print1")
	p14 := v1.MethodByName("Print4")
	fmt.Println(p11.IsValid(), p14.IsValid()) // true false

	v2 := reflect.ValueOf(&u)
	p21 := v2.MethodByName("Print1")
	p24 := v2.MethodByName("Print4")
	fmt.Println(p21.IsValid(), p24.IsValid()) // true true
}

// 反射修改值
func TestModifyValue(t *testing.T) {
	num := 1
	v1 := reflect.ValueOf(&num)
	v1.Elem().SetInt(100)
	fmt.Println("num:", num) // num: 100

	u := User{Id: 1, Name: "tom"}
	modify(&u)
	fmt.Println("user:", u) // user: {1 jerry}
}

func modify(u *User) {
	v := reflect.ValueOf(u)
	if v.Kind() != reflect.Ptr {
		fmt.Println("err: is not a pointer")
		return
	}
	if !v.Elem().CanSet() {
		fmt.Println("err: element can not set")
		return
	}

	v = v.Elem()
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("err: filed name is not valid")
		return
	}
	if f.Kind() != reflect.String {
		fmt.Println("err: field name is not string")
		return
	}
	f.SetString("jerry")
}

// 嵌套结构反射
func TestNestReflect(t *testing.T) {
	type user struct {
		Id   int
		Name string
	}

	type manger struct {
		user
		title string
	}

	m := manger{
		user: user{
			Id:   1,
			Name: "Tom",
		},
		title: "boss",
	}

	_type := reflect.TypeOf(m)
	userField, has := _type.FieldByName("user")
	fmt.Println("user field:", userField, ", has:", has)

	nameField := _type.FieldByIndex([]int{0, 0})
	fmt.Println("name field:", nameField)
}

// output:
// user field: {user go-app/ch13_reflect ch13_reflect.user  0 [0] true} , has: true
// name field: {Id  int  0 [0] false}

// 反射获取struct tag
func TestGetTag(t *testing.T) {
	type user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	u := user{Id: 1, Name: "tom"}
	_type := reflect.TypeOf(u)
	for i := 0; i < _type.NumField(); i++ {
		sf := _type.Field(i)
		tag := sf.Tag
		fmt.Println("tag:", tag)
		fmt.Println(tag.Get("json"))
	}
}

// output:
// tag: json:"id"
// id
// tag: json:"name"
// name

// 校验类型是否实现了接口
type MyError struct{}

func (*MyError) Error() string {
	return ""
}

func TestCheckImplements(t *testing.T) {
	errType := reflect.TypeOf((*error)(nil)).Elem()
	myErrType := reflect.TypeOf(MyError{})
	myErrPtrType := reflect.TypeOf(&MyError{})

	fmt.Println(myErrType.Implements(errType))    // false
	fmt.Println(myErrPtrType.Implements(errType)) // true
}

// interface -> reflection object
func TestReflectLaw1(t *testing.T) {
	str := "hello"
	fmt.Println("TypeOf str:", reflect.TypeOf(str))   // TypeOf str: string
	fmt.Println("ValueOf str:", reflect.ValueOf(str)) // ValueOf str: hello
}

// reflection object -> interface
func TestReflectLaw2(t *testing.T) {
	str := "hello"
	v := reflect.ValueOf(str)
	val := v.Interface().(string)
	fmt.Println("val:", val) // val: hello
}

// reflection object modify
func TestReflectLaw3(t *testing.T) {
	str1 := "hello"
	v1 := reflect.ValueOf(&str1)
	if v1.Kind() == reflect.Ptr && v1.Elem().CanSet() {
		v1.Elem().SetString("ok")
	}
	fmt.Println(str1) // ok

	str2 := "world"
	v2 := reflect.ValueOf(str2)
	fmt.Println(v2.Kind() == reflect.Ptr) // false
	fmt.Println(v2.CanSet())              // false
	// v2.SetString("ok")                 // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value [recovered]
}
