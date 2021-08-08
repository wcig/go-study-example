package ch32_reflet

// reflect: 运行时反射。

// 函数
// func Copy(dst, src Value) int                  // 将src的内容复制到dst中，直到dst被填满或src已用完，返回考本的元素数量，dst和src必须具有slice或array类型且具有相同的元素类型
// func DeepEqual(x, y interface{}) bool          // 报告x和y是否是深度相等
// func Swapper(slice interface{}) func(i, j int) // 返回函数用于交换给定slice的元素，slice为借口则panic

// 类型
// 1.ChanDir：表示通道类型的方向
// type ChanDir int
// func (d ChanDir) String() string
// const (
//    RecvDir ChanDir             = 1 << iota // <-chan
//    SendDir                                 // chan<-
//    BothDir = RecvDir | SendDir             // chan
// )

//  2.Kind：表示特定类型
// type Kind unit
// func (k Kind) String() string
// const (
//    Invalid Kind = iota
//    Bool
//    Int
//    Int8
//    Int16
//    Int32
//    Int64
//    Uint
//    Uint8
//    Uint16
//    Uint32
//    Uint64
//    Uintptr
//    Float32
//    Float64
//    Complex64
//    Complex128
//    Array
//    Chan
//    Func
//    Interface
//    Map
//    Ptr
//    Slice
//    String
//    Struct
//    UnsafePointer
// )

// 3.MapIter：map迭代器，参考Value.MapRange
// type MapIter struct {
//    // contains filtered or unexported fields
// }
// func (it *MapIter) Key() Value   // 返回迭代器当前map条目的key
// func (it *MapIter) Next() bool   // 报告迭代器是否还有下一个条目
// func (it *MapIter) Value() Value // 返回迭代器当前map条目的值

// 4.Method：表示单个方法
// type Method struct {
//    // Name is the method name.
//    // PkgPath is the package path that qualifies a lower case (unexported)
//    // method name. It is empty for upper case (exported) method names.
//    // The combination of PkgPath and Name uniquely identifies a method
//    // in a method set.
//    // See https://golang.org/ref/spec#Uniqueness_of_identifiers
//    Name    string
//    PkgPath string
//
//    Type  Type  // method type
//    Func  Value // func with receiver as first argument
//    Index int   // index for Type.Method
// }

// 5.SelectCase：描述选择操作的单个情况
// type SelectCase struct {
//    Dir  SelectDir // direction of case
//    Chan Value     // channel to use (for send or receive)
//    Send Value     // value to send (for send)
// }

// 6.SelectDir：描述一选择情况的同学方向
// type SelectDir int
// const (
//    SelectSend    SelectDir // case Chan <- Send
//    SelectRecv              // case <-Chan:
//    SelectDefault           // default
// )

// 7.SliceHeader：slice的运行时表示
// type SliceHeader struct {
//    Data uintptr
//    Len  int
//    Cap  int
// }

// 8.StringHeader：string的运行时表示
// type StringHeader struct {
//    Data uintptr
//    Len  int
// }

// 9.StructField：描述结构体的单个字段
// type StructField struct {
//    // Name is the field name.
//    Name string
//    // PkgPath is the package path that qualifies a lower case (unexported)
//    // field name. It is empty for upper case (exported) field names.
//    // See https://golang.org/ref/spec#Uniqueness_of_identifiers
//    PkgPath string
//
//    Type      Type      // field type
//    Tag       StructTag // field tag string
//    Offset    uintptr   // offset within struct, in bytes
//    Index     []int     // index sequence for Type.FieldByIndex
//    Anonymous bool      // is an embedded field
// }

// 10.StructTag：描述结构体字段的tag标记字符串
// type StructTag
// func (tag StructTag) Get(key string) string                     // 获取tag字符串对应key的值，没有则返回空字符串
// func (tag StructTag) Lookup(key string) (value string, ok bool) // 与Get类似，区别在于会返回key是否存在

// 11.Type：描述Go类型，为一接口。Type值可比较比如使用==运算符，也可以作为map的key。
// type Type struct {
//     ...
// }
// func ArrayOf(count int, elem Type) Type         // 返回一数组类型，该类型具有给定的count和元素类型elem
// func ChanOf(dir ChanDir, t Type) Type           // 返回一channel类型，该类型具有给定方向和元素类型的
// func FuncOf(in, out []Type, variadic bool) Type // 返回一func类型，该类型局域给定元素和返回值类型
// func MapOf(key, elem Type) Type                 // 返回一map类型，该类型具有各地的key和元素类型
// func PtrTo(t Type) Type                         // 返回给定元素t的指针类型
// func SliceOf(t Type) Type                       // 返回给定元素类型t的slice类型
// func StructOf(fields []StructField) Type        // 返回包含指定字段的struct类型
// func TypeOf(i interface{}) Type                 // 返回动态类型i的反射Type，如果i为nil的interface值则返回nil

// 12.Value：描述Go的值，类型为结构体
// type Value struct {
//     // contains filtered or unexported fields
// }
// func Append(s Value, x ...Value) Value                                 // 将值x添加到slice s中并返回结果slice
// func AppendSlice(s, t Value) Value                                     // 将slice t添加到slice s中并返回结果slice
// func Indirect(v Value) Value                                           // 返回v指向的值，如果v是nil指针则返回零值，否则返回v
// func MakeChan(typ Type, buffer int) Value                              // 创建一具有指定类型和缓冲区大小的channel
// func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value // 返回一包含指定类型typ和包装fn函数的新函数
// func MakeMap(typ Type) Value                                           // 创建具有给定类型typ的map
// func MakeMapWithSize(typ Type, n int) Value                            // 创建具有给定类型typ和初始大小n的map
// func MakeSlice(typ Type, len, cap int) Value                           // 创建具有给定类型、长度和容量的零值初始化slice
// func New(typ Type) Value                                               // 返回指向表示特定类型typ的新零值的指针的Value，即返回Value的类型为PtrTo(typ)
// func NewAt(typ Type, p unsafe.Pointer) Value                           // 返回一个表示指向指定类型值的指针的 Value，使用 p 作为该指针。
// func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)  // 指向列表描述的选择操作
// func ValueOf(i interface{}) Value                                      // 返回一新的Value，其初始化为interface i存储具体值，i为nil则返回零值
// func Zero(typ Type) Value                                              // 返回指定类型的零值Value
// func (v Value) Addr() Value // 返回一表示v地址的指针值
// func (v Value) Bool() bool // 返回v的底层bool值，类型为非bool则panic
// func (v Value) Bytes() []byte // 返回v的底层[]byte值，类型不对则panic
// func (v Value) Call(in []Value) []Value // 使用指定入参in调用函数v
// func (v Value) CallSlice(in []Value) []Value // 使用指定入参in调用可变函数v
// func (v Value) CanAddr() bool // 报告是否可通过Addr获取value的地址
// func (v Value) CanInterface() bool // 报告interface是否可使用而不panic
// func (v Value) CanSet() bool // 报告v的值是否可被修改
// func (v Value) Cap() int // 返回v的容量，类型不是Array、Chan、Slice则panic
// func (v Value) Close() // 关闭通道v，类型不是Chan则panic
// func (v Value) Complex() complex128 // 返回v的底层complex128类型值
// func (v Value) Convert(t Type) Value // 返回v转换为类型t的值
// func (v Value) Elem() Value // 返回接口v包含的值或指针v指向的值，非接口或指针则panic
// func (v Value) Field(i int) Value // 返回结构体v的第i个字段，类型非结构体或i越界则panic
// func (v Value) FieldByIndex(index []int) Value // 返回索引index对应的嵌套字段（用于嵌套结构体）
// func (v Value) FieldByName(name string) Value // 获取结构体中指定name字段的值
// func (v Value) FieldByNameFunc(match func(string) bool) Value // 返回结构体字段中满足匹配函数match的值
// func (v Value) Float() float64 // 返回v底层float64值（v类型非float32、float64则panic）
// func (v Value) Index(i int) Value 返回v的第i个元素值（v类型非Array、Slice、String或越界则panic）
// func (v Value) Int() int64 // 返回v的底层int值（v类型非int、int8、int16、int32、int64则panic）
// func (v Value) Interface() (i interface{}) // 返回v当前值作为interface{}
// func (v Value) InterfaceData() [2]uintptr // 返回接口v的值作为uintptr对
// func (v Value) IsNil() bool // 报告v是否为nil
// func (v Value) IsValid() bool // 报告v是否表示一值，如果v为零值则返回false
// func (v Value) IsZero() bool // 报告v是否为其类型的零值
// func (v Value) Kind() Kind // 返回v的类型Kind
// func (v Value) Len() int // 返回v的长度（v类型为非Array、Chan、Map、Slice或String则panic）
// func (v Value) MapIndex(key Value) Value // 返回map v的key对应值
// func (v Value) MapKeys() []Value // 返回一slice包含map所有key对应的值
// func (v Value) MapRange() *MapIter // 返回map的范围迭代器
// func (v Value) Method(i int) Value // 返回v的第i个方法对应函数值
// func (v Value) MethodByName(name string) Value // 返回v的方法名为name对应函数值
// func (v Value) NumField() int // 返回结构体v的字段个数
// func (v Value) NumMethod() int // 返回可导出方法的个数
// func (v Value) OverflowComplex(x complex128) bool // 报告complex128 x是否不能用v的类型表示。如果 v 的 Kind 不是 Complex64 或 Complex128，它会发生恐慌。
// func (v Value) OverflowFloat(x float64) bool // 报告是否float64 x的不能用v的类型表示
// func (v Value) OverflowInt(x int64) bool // 报告是否int64 x的不能用v的类型表示
// func (v Value) OverflowUint(x uint64) bool // 报告是否uint64 x的不能用v的类型表示
// func (v Value) Pointer() uintptr // 返回v的uintptr的值
// func (v Value) Recv() (x Value, ok bool) // 从通道v接收并返回一个值
// func (v Value) Send(x Value) // 在通道v发送x
// func (v Value) Set(x Value) // 分配x给到值v
// func (v Value) SetBool(x bool) // 设置v的底层bool类型值
// func (v Value) SetBytes(x []byte) // 设置v的底层[]byte类型值
// func (v Value) SetCap(n int) // 设置v的容量为n
// func (v Value) SetComplex(x complex128) // 设置v的底层complex128类型值
// func (v Value) SetFloat(x float64) // 设置v的底层float64类型值
// func (v Value) SetInt(x int64) // 设置v的底层int64类型值
// func (v Value) SetLen(n int) // 设置v的长度为n
// func (v Value) SetMapIndex(key, elem Value) // 设置map v的key和值
// func (v Value) SetPointer(x unsafe.Pointer) // 设置unsafe.Pointer值v为x
// func (v Value) SetString(x string) // 设置v的底层string类型值
// func (v Value) SetUint(x uint64) // 设置v的底层uint64类型值
// func (v Value) Slice(i, j int) Value // 返回slice v[i:j]
// func (v Value) Slice3(i, j, k int) Value // 返回slice v[i:j:k]
// func (v Value) String() string // 返回v底层值的字符串值
// func (v Value) TryRecv() (x Value, ok bool) // 以非阻塞方式从通道v接收值
// func (v Value) TrySend(x Value) bool // 以非阻塞方式向通道v发送x
// func (v Value) Type() Type // 返回v的类型
// func (v Value) Uint() uint64 // 返回v的底层uint值
// func (v Value) UnsafeAddr() uintptr // 返回指向v数据的指针

// 13.ValueError：类型错误，比如在不支持的值上调用方法
// type ValueError struct {
//    Method string
//    Kind   Kind
// }
// func (e *ValueError) Error() string
