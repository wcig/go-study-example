package ch25_log

import (
	"fmt"
	"log"
	"testing"
)

// log: 简单的日志包实现。

// 常量
func TestConst(t *testing.T) {
	// 日志格式
	_ = log.Ldate         // the date in the local time zone: 2009/01/23
	_ = log.Ltime         // the time in the local time zone: 01:23:23
	_ = log.Lmicroseconds // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	_ = log.Llongfile     // full file name and line number: /a/b/c/d.go:23
	_ = log.Lshortfile    // final file name element and line number: d.go:23. overrides Llongfile
	_ = log.LUTC          // if Ldate or Ltime is set, use UTC rather than the local time zone
	_ = log.Lmsgprefix    // move the "prefix" from the beginning of the line to before the message
	_ = log.LstdFlags     // initial values for the standard logger
}

// 函数
// 写日志
// 1.log.Fatal[f|ln]: 写入日志消息并调用os.Exit(1)
// func Fatal(v ...interface{})
// func Fatalf(format string, v ...interface{})
// func Fatalln(v ...interface{})

// 2.log.Panic[f|ln]: 写入日志消息并调用panic
// func Panic(v ...interface{})
// func Panicf(format string, v ...interface{})
// func Panicln(v ...interface{})

// 3.log.Print[f|ln]: 只写入日志消息
// func Print(v ...interface{})
// func Printf(format string, v ...interface{})
// func Println(v ...interface{})

// 其他
// func Flags() int                           // 返回标准logger的输出标志，比如Ldata、Ltime等
// func Output(calldepth int, s string) error // 日志写入s，同时calldepth表示跳过函数调用的层级，值为1标识打印输出调用者详细信息
// func Prefix() string                       // 返回标准logger的输出前缀
// func SetFlags(flag int)                    // 设置标准logger的输出标志，比如Ldata、Ltime等
// func SetOutput(w io.Writer)                // 设置标准logger的输出目的地
// func SetPrefix(prefix string)              // 设置标准logger的输出前缀
// func Writer() io.Writer                    // 返回标准logger的输出目的地

// 类型
// type log.Logger: 一活动的日志对象，写入输出到io.Writer，一个Logger可被多个goroutines同时使用，它可以保证对Writer的序列化访问。
// func Default() *Logger                                   // 返回默认包级别Logger
// func New(out io.Writer, prefix string, flag int) *Logger // 基于输出、前缀和标志创建一新的Logger
// func (l *Logger) Fatal(v ...interface{})                 // ...
// func (l *Logger) Fatalf(format string, v ...interface{})
// func (l *Logger) Fatalln(v ...interface{})
// func (l *Logger) Flags() int
// func (l *Logger) Output(calldepth int, s string) error
// func (l *Logger) Panic(v ...interface{})
// func (l *Logger) Panicf(format string, v ...interface{})
// func (l *Logger) Panicln(v ...interface{})
// func (l *Logger) Prefix() string
// func (l *Logger) Print(v ...interface{})
// func (l *Logger) Printf(format string, v ...interface{})
// func (l *Logger) Println(v ...interface{})
// func (l *Logger) SetFlags(flag int)
// func (l *Logger) SetOutput(w io.Writer)
// func (l *Logger) SetPrefix(prefix string)
// func (l *Logger) Writer() io.Writer
func TestTypeLogger(t *testing.T) {
	_ = log.Logger{}
	// type Logger struct {
	//    // contains filtered or unexported fields
	// }
}

func TestDefaulLogger(t *testing.T) {
	logger := log.Default()

	fmt.Println(logger.Flags())
	fmt.Println(logger.Prefix())

	logger.SetPrefix("mylog: ")
	logger.SetFlags(log.Lshortfile | logger.Flags())
	fmt.Println(logger.Prefix())
	logger.Println("ok")

	infof := func(info string) {
		logger.Output(1, info)
	}
	infof("hello world.")
	// output:
	// 3
	//
	// mylog:
	// mylog: 2021/06/18 22:34:21 log_test.go:86: ok
	// mylog: 2021/06/18 22:34:21 log_test.go:89: hello world.
}
