
> go 代码规范

#### 1. 命名/import
##### 1.1 包名
包名使用小写，使用短命名，不与标准库冲突。
包名应该用单数的形式，比如util、model,而不是utils、models。

##### 1.2 常量命名
常量命名推荐使用首字母大写的驼峰格式。
eg：使用 MaxLength ，而不是 MAX_LENGTH、MAXLENGTH

##### 1.3 变量命名
全局变量使用首字母大写的驼峰格式，局部变量使用首字母小写的驼峰格式。
eg：全局变量：MaxLenght，局部变量：maxLength

##### 1.4 命名保证统一
对于缩写词：URL、HTTP、ID这种，或者都大写，或者都小写。

##### 1.5 多个变量声明方式
函数中当需要使用到多个变量时，可以在函数开始处使用var声明。
在函数外部申明必须使用var,不要采用:=，容易踩到变量的作用域的问题。
```go
var (
	Width  int
	Height int
)
```

##### 1.6 import规范
import在多行的情况下，goimports会自动帮你格式化，但是我们这里还是规范一下import的一些规范，如果你在一个文件里面引入了一个package，还是建议采用如下格式：
```go
import (
    "fmt"
)
```
如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包：
```go
import (
    "encoding/json"
    "strings"

    "myproject/models"
    "myproject/controller"
    "myproject/utils"

    "github.com/astaxie/beego"
    "github.com/go-sql-driver/mysql"
)  
```
在项目中不要使用相对路径引入包：
```go
// 推荐
import "github.com/repo/proj/src/net"

// 不推荐
import "../net"
```

---


#### 2.字符串
##### 2.1 空字符串判断
```go
// 推荐
if s == "" {
    ...
}

// 不推荐
if len(s) == 0 {
    ...
}

// 语法错误(s==nil)
if s == nil || len(s) == 0 {
    ...
}
```

##### 2.2 []byte/string相等比较
```go
// 推荐
var s1 []byte
var s2 []byte
...
bytes.Equal(s1, s2) == 0
bytes.Equal(s1, s2) != 0

// 不推荐
var s1 []byte
var s2 []byte
...
bytes.Compare(s1, s2) == 0
bytes.Compare(s1, s2) != 0
```

##### 2.3 字符串是否包含子串或字符
```go
// 推荐
strings.Contains(s, subStr)
strings.ContainsAny(s, char)
strings.ContainRune(s, r)

// 不推荐
strings.Index(s, subStr) > -1
strings.IndexAny(s, char) > -1
strings.IndexRune(s, r) > -1
```

##### 2.4 去除前后子串
```go
// 推荐
var s1 = "a string value"
var s2 = "a "
var s3 = strings.TrimPrefix(s1, s2)

// 不推荐
var s1 = "a string value"
var s2 = "a "
var s3 string
if strings.HasPrefix(s1, s2) {
    s3 = s1[len(s2):]
}
```

##### 2.5 复杂字符串使用raw字符串避免字符转义
```go
//不推荐
regexp.MustCompile("\\.")

//推荐
regexp.MustCompile(`\.`)
```

---

#### 3. slice
##### 3.1 空slice判断
```go
// 推荐
if len(slice) > 0 {
    ...
}

// 不推荐
if slice != nil && len(slice) > 0 {
    ...
}
```
上面判断同样适用于map、channel

##### 3.2 声明slice
```go
// 建议声明方式：
var s []string

// 不建议声明方式：
s := []string{}
s := make([]string, 0)
```
原因：
前者声明了一个nil slice，后两者是定义了一个长度为0的非nil slice。

##### 3.3 slice复制
```go
// 推荐
copy(b2, b1)

// 不推荐
var b1, b2 []byte
for i, v := range b1 {
   b2[i] = v
}
for i := range b1 {
   b2[i] = b1[i]
}
```

##### 3.4 slice新增
```go
// 推荐
var a, b []int
b = append(b, a...)

// 不推荐
var a, b []int
for _, v := range a {
    b = append(b, v)
}
```

---


#### 4. struct
##### 4.1 struct初始化
struct以多行格式初始化
```go
type user struct {
	Id   int64
	Name string
}

u1 := user{10001, "Nick"}

u2 := user{
    Id:   10002,
    Name: "Tom",
}
```

##### 4.2 struct方法
Receiver的名称建议以缩写形式，简短的使用一个或两个字符，service类使用service。
```go
type foo struct {}

func (f foo)method() {
	//...
}

type userService struct {}

func (service *userService)method() {
	//...
}
```

##### 4.3 struct转换
如果两个struct数据类型一致，可直接转换。
如果两个struct数据类型一致，但json tag不一致，也可直接转换。
```go
type T1 struct {
	A int
	B int
}
type T2 struct {
	A int
	B int
}
func TestTransStruct(t *testing.T) {
	t1 := T1{
		A: 1,
		B: 2,
	}
	t2 := T2{
		A: t1.A,
		B: t1.B,
	}
	fmt.Println("t2:", t2)

	//可直接转换
	t3 := T2(t1)
	fmt.Println("t3:", t3)
}
```

---


#### 5. 错误处理
##### 5.1 error、panic
error：不要使用 _ 忽略错误，保证处理每一个错误。
panic：尽量避免使用panic。只有在不可运行的情况采用panic，例如文件无法打开，数据库无法连接导致程序无法正常运行，但是对于其他的package对外的接口不能有panic，只能在包内采用。

##### 5.2 错误信息
建议把自定义的Error放在package级别中，统一进行维护。error命名以Err开头，同时error信息不要使用大写字母，足够简短又能表达错误信息。
```go
var (
    ErrCacheMiss = errors.New("memcache: cache miss")
    ErrServerError = errors.New("memcache: server error")
)
```

---


#### 6. 其他
##### 6.1 bool判断
```go
// 推荐
if flag {
    ...
}

// 不推荐
if flag == true {
    ...
}
```

##### 6.2 for循环
```go
// 推荐
for {
    ...
}

// 不推荐
for true {
    ...
}
```

##### 6.3 简化range
```go
var m map[string]int
for _, val := range m {
    ...
}

for index, _ := range m {
    ...
}
```

##### 6.4 省略不必要的变量
通过 _ 省略不必要的变量。

##### 6.5 通过if-return减少if分支嵌套(嵌套建议不超过3层)
```go
// 推荐
if {
	return
}
if {
	...
}

// 不推荐
if {
	if {
		...
	}
}
```

##### 6.6 time.Since
```go
tt := time.Date(2019, 5, 25, 0, 0, 0, 0, time.Local)
t1 := time.Now().Sub(tt)
t2 := time.Since(tt)
```

##### 6.7 函数多返回值
函数有多返回值时，建议函数定义处为返回值命名，这样函数定义更清晰。
```go
// 推荐
func getSize(videoDir string) (width, height int) {
	//...
}

// 不推荐
func getSize(videoDir string) (int, int) {
	//...
}
```

参考：     
1. [Go编码规范指南](https://gocn.vip/article/1)      
2. [编写地道的Go代码](https://colobu.com/2017/02/07/write-idiomatic-golang-codes/)        
3. [Golang代码规范](https://sheepbao.github.io/post/golang_code_specification/)