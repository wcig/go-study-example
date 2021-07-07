package exec

import (
	"os/exec"
	"testing"
)

// os/exec: 运行外部命令，包装了os.StartProcess以便更方便的重新映射stdin和stdout、用管道连接I/O和做其他调整。

// 错误
func TestErr(t *testing.T) {
	_ = exec.ErrNotFound // executable file not found in $PATH
}

// 函数
// func LookPath(file string) (string, error) // 在PATH环境变量命名的目录中搜索一名称为file的可执行文件

// 类型
// 1.Cmd: 表示正准备和运行的 外部命令，在调用其Run、Output或CombinedOutpu方法后，不能重复该Cmd
// type Cmd struct {
//    Path string
//    Args []string
//    Env []string
//    Dir string
//    Stdin io.Reader
//    Stdout io.Writer
//    Stderr io.Writer
//    ExtraFiles []*os.File
//    SysProcAttr *syscall.SysProcAttr
//    Process *os.Process
//    ProcessState *os.ProcessState
//    // contains filtered or unexported fields
// }
// func Command(name string, arg ...string) *Cmd                             // 以命令name额参数arg创建一Cmd
// func CommandContext(ctx context.Context, name string, arg ...string) *Cmd // 与Command类型，但包含了一个Context，如果context在命令自行完成之前完成，则context用于终止进程
// func (c *Cmd) CombinedOutput() ([]byte, error)                            // 运行命令并返回标准输出和标准错误的组合
// func (c *Cmd) Output() ([]byte, error)                                    // 运行命令，返回标准输出，任何返回错误通常是*ExitError类型，如果c.Stderr为nil，输出为ExitError.Stderr
// func (c *Cmd) Run() error                                                 // 运行命令并等待完成，如果正确运行返回nil，如果命令未完成返回类型*ExitError错误，其他情况可能返回其他错误类型。
// func (c *Cmd) Start() error                                               // 开始进程但不等待其完成，如果返回成功则将设置c.Process字段
// func (c *Cmd) StderrPipe() (io.ReadCloser, error)                         // 返回一管道，该管道在命令启动时连接命令的标准错误
// func (c *Cmd) StdinPipe() (io.WriteCloser, error)                         // 返回一管道，用于在命令启动时连接命令的标准输入
// func (c *Cmd) StdoutPipe() (io.ReadCloser, error)                         // 返回一管道，用于在命令启动时连接命令的标准输出
// func (c *Cmd) String() string                                             // 返回用户可读的c描述，仅用于调试
// func (c *Cmd) Wait() error                                                // 等待命令退出并等待任何复制到stdin或stdout或stderr复制完成，改名了必须由Start启动

// 2.Error: 当LookPath未能将文件分类为可执行文件返回此错误
// type Error struct {
//    // Name is the file name for which the error occurred.
//    Name string
//    // Err is the underlying error.
//    Err error
// }
// func (e *Error) Error() string
// func (e *Error) Unwrap() error

// 3.ExitError: 报告命令非成功退出
// type ExitError struct {
//    *os.ProcessState
//    Stderr []byte // Go 1.6
// }
//    func (e *ExitError) Error() string
