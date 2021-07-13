package automic

// sync/automic: 提供低级的原子操作工具函数

// 函数
// // 原子交换
// func SwapInt32(addr *int32, new int32) (old int32)
// func SwapInt64(addr *int64, new int64) (old int64)
// func SwapUint32(addr *uint32, new uint32) (old uint32)
// func SwapUint64(addr *uint64, new uint64) (old uint64)
// func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
// func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
//
// // 原子比较并交换（只有当操作之前的值为old才进行交换操作）
// func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
// func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
// func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
// func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
// func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
// func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
//
// // 原子修改值
// func AddInt32(addr *int32, delta int32) (new int32)
// func AddUint32(addr *uint32, delta uint32) (new uint32)
// func AddInt64(addr *int64, delta int64) (new int64)
// func AddUint64(addr *uint64, delta uint64) (new uint64)
// func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
//
// // 原子获取值
// func LoadInt32(addr *int32) (val int32)
// func LoadInt64(addr *int64) (val int64)
// func LoadUint32(addr *uint32) (val uint32)
// func LoadUint64(addr *uint64) (val uint64)
// func LoadUintptr(addr *uintptr) (val uintptr)
// func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
//
// // 原子设定值
// func StoreInt32(addr *int32, val int32)
// func StoreInt64(addr *int64, val int64)
// func StoreUint32(addr *uint32, val uint32)
// func StoreUint64(addr *uint64, val uint64)
// func StoreUintptr(addr *uintptr, val uintptr)
// func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

// 类型
// Value: 提供原子的加载和存储类型一致的值. 零值Load时返回nil,一旦Store被调用该Value不能被复制. (在第一次使用后不能复制)
// type Value struct {
//    // contains filtered or unexported fields
// }
// func (v *Value) Load() (x interface{})
// func (v *Value) Store(x interface{})
