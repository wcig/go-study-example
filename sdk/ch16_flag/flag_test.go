package main

import (
	"flag"
	"testing"
)

// flag: 命令行标志解析
// flag 包支持的命令行参数类型有 bool、int、int64、uint、uint64、float、float64、string、duration.

// 命令行语法：
// -flag
// -flag=x
// -flag x // non-boolean flags only

// 变量
func TestVar(t *testing.T) {
	_ = flag.CommandLine // 默认解析命令行flag集合
}

// 错误
func TestErr(t *testing.T) {
	_ = flag.ContinueOnError // 描述性错误
	_ = flag.ExitOnError     // 调用os.Exit(2) 或-h/-help Exit(0)
	_ = flag.PanicOnError    // 描述性的panic错误
}

// 函数
// 1.命令行参数解析方式一: flag.Type() *Type
// func Bool(name string, value bool, usage string) *bool
// func Duration(name string, value time.Duration, usage string) *time.Duration
// func Float64(name string, value float64, usage string) *float64
// func Int(name string, value int, usage string) *int
// func Int64(name string, value int64, usage string) *int64
// func String(name string, value string, usage string) *string
// func Uint(name string, value uint, usage string) *uint
// func Uint64(name string, value uint64, usage string) *uint64

// 2.命令行参数解析方式二: flag.TypeVar()
// func BoolVar(p *bool, name string, value bool, usage string)
// func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
// func Float64Var(p *float64, name string, value float64, usage string)
// func IntVar(p *int, name string, value int, usage string)
// func Int64Var(p *int64, name string, value int64, usage string)
// func StringVar(p *string, name string, value string, usage string)
// func UintVar(p *uint, name string, value uint, usage string)
// func Uint64Var(p *uint64, name string, value uint64, usage string)

// 3.其他
// func Arg(i int) string // 返回命令行参数中,第i个非标志参数(从0开始),没有则返回空字符串
// func Args() []string   // 返回命令行参数中所有的非标志参数
// func NArg() int // 已处理标志后剩余的参数的数量
// func NFlag() int // 已设置的命令行标志数
// func Parse() // 解析来自OS.ARGS [1：]的命令行标志,定义标志后必须调用
// func Parsed() bool // 报告命令行标志是否已解析
// func Set(name, value string) error // 设置命名命令行标志的值

// 4.结构体：定义flag的集合
// type FlagSet struct {
// 	Usage func()
// 	// 其他未导出字段
// }
