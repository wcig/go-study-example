package ch29_os

// os：提供了操作系统功能的独立于平台的接口，设计类Unix。

// 1.常量
// Flags：包装底层系统标志
// const (
// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// 	// The remaining values may be or'ed in to control behavior.
// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
// )

// 分隔符
// const (
// 	PathSeparator     = '/' // OS-specific path separator
// 	PathListSeparator = ':' // OS-specific path list separator
// )

// 文件模式位
// const (
// 	// The single letters are the abbreviations
// 	// used by the String method's formatting.
// 	ModeDir        = fs.ModeDir        // d: is a directory
// 	ModeAppend     = fs.ModeAppend     // a: append-only
// 	ModeExclusive  = fs.ModeExclusive  // l: exclusive use
// 	ModeTemporary  = fs.ModeTemporary  // T: temporary file; Plan 9 only
// 	ModeSymlink    = fs.ModeSymlink    // L: symbolic link
// 	ModeDevice     = fs.ModeDevice     // D: device file
// 	ModeNamedPipe  = fs.ModeNamedPipe  // p: named pipe (FIFO)
// 	ModeSocket     = fs.ModeSocket     // S: Unix domain socket
// 	ModeSetuid     = fs.ModeSetuid     // u: setuid
// 	ModeSetgid     = fs.ModeSetgid     // g: setgid
// 	ModeCharDevice = fs.ModeCharDevice // c: Unix character device, when ModeDevice is set
// 	ModeSticky     = fs.ModeSticky     // t: sticky
// 	ModeIrregular  = fs.ModeIrregular  // ?: non-regular file; nothing else is known about this file
//
// 	// Mask for the type bits. For regular files, none will be set.
// 	ModeType = fs.ModeType
//
// 	ModePerm = fs.ModePerm // Unix permission bits, 0o777
// )

// 空设备
// const DevNull = "/dev/null"

// 2.变量
// var (
// 	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// 	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// 	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
// )

// 3.错误
// var (
// 	// ErrInvalid indicates an invalid argument.
// 	// Methods on File will return this error when the receiver is nil.
// 	ErrInvalid = fs.ErrInvalid // "invalid argument"
//
// 	ErrPermission = fs.ErrPermission // "permission denied"
// 	ErrExist      = fs.ErrExist      // "file already exists"
// 	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
// 	ErrClosed     = fs.ErrClosed     // "file already closed"
//
// 	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"
// 	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
// )
// var ErrProcessDone = errors.New("os: process already finished")

// 4.函数
// func Chdir(dir string) error                                      // 创建目录
// func Chmod(name string, mode FileMode) error                      // 修改name文件权限
// func Chown(name string, uid, gid int) error                       // 修改name文件所属用户和所属用户组
// func Chtimes(name string, atime time.Time, mtime time.Time) error // 修改name文件的访问时间和修改时间
// func Clearenv()                                                   // 删除所有环节变量
// func DirFS(dir string) fs.FS                                      // 以dir为根的文件数返回一文件系统fs.FS
// func Environ() []string                                           // 以"key=value"返回环节变量
// func Executable() (string, error)                                 // 返回启动当前进程的可执行文件的路径名
// func Exit(code int) // 以指定状态码code退出当前程序，0表示成功，非0表示错误。程序会立即终止，defer函数不会运行（为了可移植性，状态码应在[0,125]范围内）
// func Expand(s string, mapping func(string) string) string // 根据mapping映射函数替换字符串s中的${var}或$var
// func ExpandEnv(s string) string // 根据环境变量替换字符串s中的${var}或$var
// func Getegid() int // 返回调用者的有效数字组id，windows返回-1
// func Getenv(key string) string // 获取环境变量中指定key对应的值，不存在则返回空
// func Geteuid() int // 返回调用者的有效数字用户id，windows返回-1
// func Getgid() int // 返回调用者的数字组id，windows返回-1
// func Getgroups() ([]int, error)
// func Getpagesize() int
// func Getpid() int
// func Getppid() int
// func Getuid() int
// func Getwd() (dir string, err error)
// func Hostname() (name string, err error)
// func IsExist(err error) bool
// func IsNotExist(err error) bool
// func IsPathSeparator(c uint8) bool
// func IsPermission(err error) bool
// func IsTimeout(err error) bool
// func Lchown(name string, uid, gid int) error
// func Link(oldname, newname string) error
// func LookupEnv(key string) (string, bool)
// func Mkdir(name string, perm FileMode) error
// func MkdirAll(path string, perm FileMode) error
// func MkdirTemp(dir, pattern string) (string, error)
// func NewSyscallError(syscall string, err error) error
// func Pipe() (r *File, w *File, err error)
// func ReadFile(name string) ([]byte, error)
// func Readlink(name string) (string, error)
// func Remove(name string) error
// func RemoveAll(path string) error
// func Rename(oldpath, newpath string) error
// func SameFile(fi1, fi2 FileInfo) bool
// func Setenv(key, value string) error
// func Symlink(oldname, newname string) error
// func TempDir() string
// func Truncate(name string, size int64) error
// func Unsetenv(key string) error
// func UserCacheDir() (string, error)
// func UserConfigDir() (string, error)
// func UserHomeDir() (string, error)
// func WriteFile(name string, data []byte, perm FileMode) error
