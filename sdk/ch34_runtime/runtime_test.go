package ch34_runtime

import (
	"runtime"
	"testing"
)

// runtime: 包含与 Go 的运行时系统交互的操作，例如控制 goroutine 的函数。

// 常量
func TestConst(t *testing.T) {
	_ = runtime.Compiler // gc: 当前构建运行的二进制文件编译器工具链
	_ = runtime.GOARCH   // amd64: 程序运行架构（386、amd64、arm、s390x等）
	_ = runtime.GOOS     // darwin: 程序运行操作系统（drawin、freebsd、linux等）
}

// 变量
func TestVar(t *testing.T) {
	_ = runtime.MemProfileRate // 512*1024
}

// 函数
// func GC()                  // 执行一次垃圾回收GC并阻塞调用者，直到GC完成，可能阻塞整个程序。
// func GOMAXPROCS(n int) int // 设置可以同时执行的最大CPU个数，并返回之前的设置，默认值为runtime.NumCPu。
// func GOROOT() string       // 返回Go树的根
// func Goexit()              // 终止调用它的goroutine，其他goroutine不受影响。在终止goroutine前会执行所有的defer语句，因为Goexit不是panic，所以所有defer函数recover都会返回nil。
// func Gosched()             // 让出处理器，允许其他goroutine允许。它不会挂起当前goroutine，锁执行会自动恢复
// func NumCPU()              // 返回当前进程可使用的逻辑CPU个数
// func NumCgoCall()          // 返回当前进程调用的cgo次数
// func NumGoroutine()        // 返回当前存在的goroutine数量
// func ReadMemStats(m *MemStats) // 设置内存分配器统计信息到m
// func Version() string       // 返回Go版本

// 类型
// 1.Error: 定义运行时错误
// type Error interface {
// 	error
// 	RuntimeError()
// }

// 2.MemStats: 内存分配器统计信息
// type MemStats struct{
// 	...
// }

// 子包
// cgo: cgo 包含对 cgo 工具生成的代码的运行时支持。
// debug: debug 包含程序在运行时自行调试的工具。
// metrics: 指标提供了一个稳定的接口来访问由 Go 运行时导出的实现定义的指标。
// msan:
// pprof: pprof 以 pprof 可视化工具预期的格式写入运行时分析数据。
// Race: Race 实现了数据竞争检测逻辑。
// trace: 跟踪包含程序为 Go 执行跟踪器生成跟踪的工具。
