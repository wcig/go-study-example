package fs

import (
	"io/fs"
	"testing"
)

// io/fs：定义了文件系统的基本接口。文件系统可以由主机操作系统提供，也可以由其他包提供。

// 错误
func TestErr(t *testing.T) {
	_ = fs.ErrInvalid    // "invalid argument"
	_ = fs.ErrPermission // "permission denied"
	_ = fs.ErrExist      // "file already exists"
	_ = fs.ErrNotExist   // "file does not exist"
	_ = fs.ErrClosed     // "file already closed"

	_ = fs.SkipDir // 用作 WalkDirFuncs 的返回值，以指示要跳过调用中指定的目录。它不会被任何函数作为错误返回。
}

// 函数
// func Glob(fsys FS, pattern string) (matches []string, err error) // 返回所有匹配pattern的文件名称，没有则返回nil
// func ReadFile(fsys FS, name string) ([]byte, error)              // 读取文件系统fs指定name的文件，返回其内容，成功调用err==nil而不是EOF
// func ValidPath(name string) bool                                 // 校验指定路径名在调用Open方法时是否有效
// func WalkDir(fsys FS, root string, fn WalkDirFunc) error         // 遍历以root为根的文件数，为树的每个文件或目录包括root调用fn方法。
// func ReadDir(fsys FS, name string) ([]DirEntry, error) // 读取目录name，返回按文件名排序的目录条目列表
// func Sub(fsys FS, dir string) (FS, error) // 返回以fsys为根目录的目录为dir的FS
// func Stat(fsys FS, name string) (FileInfo, error) // 返回fsys文件系统指定name文件的文件信息

// 类型
// 1.type DirEntity: 从目录中读取的条目（使用ReadDir函数或ReadDirFile的ReadDir方法）
// type DirEntry interface {
//    Name() string
//    IsDir() bool
//    Type() FileMode
//    Info() (FileInfo, error)
// }

// 2.type FS: 提供对分层系统的访问
// type FS interface {
//    Open(name string) (File, error)
// }

// 3.type File: 提供对单个文件的访问
// type File interface {
//    Stat() (FileInfo, error)
//    Read([]byte) (int, error)
//    Close() error
// }

// 4.type FileInfo: 用于描述文件，调用Stat返回
// type FileInfo interface {
//    Name() string       // base name of the file
//    Size() int64        // length in bytes for regular files; system-dependent for others
//    Mode() FileMode     // file mode bits
//    ModTime() time.Time // modification time
//    IsDir() bool        // abbreviation for Mode().IsDir()
//    Sys() interface{}   // underlying data source (can return nil)
// }

// 5.type FileMode: 文件模式和权限位
// type FileMode uint32
// const (
//    // The single letters are the abbreviations
//    // used by the String method's formatting.
//    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
//    ModeAppend                                     // a: append-only
//    ModeExclusive                                  // l: exclusive use
//    ModeTemporary                                  // T: temporary file; Plan 9 only
//    ModeSymlink                                    // L: symbolic link
//    ModeDevice                                     // D: device file
//    ModeNamedPipe                                  // p: named pipe (FIFO)
//    ModeSocket                                     // S: Unix domain socket
//    ModeSetuid                                     // u: setuid
//    ModeSetgid                                     // g: setgid
//    ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
//    ModeSticky                                     // t: sticky
//    ModeIrregular                                  // ?: non-regular file; nothing else is known about this file
//
//    // Mask for the type bits. For regular files, none will be set.
//    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular
//
//    ModePerm FileMode = 0777 // Unix permission bits
// )
// 方法：
// func (m FileMode) IsDir() bool: 报告m是否目录
// func (m FileMode) IsRegular() bool: 报告m是都为一常规文件
// func (m FileMode) Perm() FileMode: 返回m的unix权限位
// func (m FileMode) String() string: m信息字符串输出
// func (m FileMode) Type() FileMode: 返回m比特位

// 6.type GlobFS: Glob方法的文件系统
// type GlobFS interface {
// 	FS
// 	Glob(pattern string) ([]string, error)
// }

// 7.type PathError: 记录错误和操作以及导致它的文件路径
// type PathError struct {
//    Op   string
//    Path string
//    Err  error
// }
// 方法：
// func (e *PathError) Error() string // 字符串输出
// func (e *PathError) Timeout() bool // 报告该错误是否表示超时
// func (e *PathError) Unwrap() error

// 8.type ReadDirFS: 文件系统的接口实现，它提供了ReadDir的优化实现
// type ReadDirFS interface {
//    FS
//    ReadDir(name string) ([]DirEntry, error)
// }

// 9.type ReadDirFile: 一目录文件，包括ReadDir方法
// type ReadDirFile interface {
// 	File
// 	ReadDir(n int) ([]DirEntry, error)
// }

// 10.type ReadFileFS: 文件系统实现，提供了ReadFile的优化实现
// type ReadFileFS interface {
// 	File
// 	ReadFile(name string) ([]byte, error)
// }

// 11.type StatFS: 包含Stat方法的文件系统实现
// type StatFS interface {
// 	FS
// 	Stat(name string) (FileInfo, error)
// }

// 12.type SubFS: 包含Sub方法的文件系统实现
// type SubFS interface {
// 	FS
// 	Sub(dir string) (FS, error)
// }

// 13.type WalkDirFunc: WalkDir调用的函数类型，用于访问每个文件或目录。
// type WalkDirFunc func(path string, d DirEntry, err error) error
