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
// func Chdir(dir string) error                                      // 修改当前工作目录为name目录
// func Chmod(name string, mode FileMode) error                      // 修改name文件权限
// func Chown(name string, uid, gid int) error                       // 修改name文件所属用户和所属用户组
// func Chtimes(name string, atime time.Time, mtime time.Time) error // 修改name文件的访问时间和修改时间
// func Clearenv()                                                   // 删除所有环节变量
// func DirFS(dir string) fs.FS                                      // 以dir为根的文件数返回一文件系统fs.FS
// func Environ() []string                                           // 以"key=value"返回环节变量
// func Executable() (string, error)                                 // 返回启动当前进程的可执行文件的路径名
// func Exit(code int)                                               // 以指定状态码code退出当前程序，0表示成功，非0表示错误。程序会立即终止，defer函数不会运行（为了可移植性，状态码应在[0,125]范围内）
// func Expand(s string, mapping func(string) string) string         // 根据mapping映射函数替换字符串s中的${var}或$var
// func ExpandEnv(s string) string                                   // 根据环境变量替换字符串s中的${var}或$var
// func Getegid() int                                                // 返回调用者的有效数字组id，windows返回-1
// func Getenv(key string) string                                    // 获取环境变量中指定key对应的值，不存在则返回空
// func Geteuid() int                                                // 返回调用者的有效数字用户id，windows返回-1
// func Getgid() int                                                 // 返回调用者的数字组id，windows返回-1
// func Getgroups() ([]int, error)                                   // 返回调用者所属组的数字id列表
// func Getpagesize() int                                            // 返回底层系统的内存页大小
// func Getpid() int                                                 // 返回调用者的进程id
// func Getppid() int                                                // 返回调用者的父进程id
// func Getuid() int                                                 // 返回调用者的用户数字id
// func Getwd() (dir string, err error)                              // 返回当前目录的根路径名
// func Hostname() (name string, err error)                          // 返回内核报告的主机名
// func IsExist(err error) bool                                      // 报告错误是否为文件或目录已存在（ErrExist满足）
// func IsNotExist(err error) bool                                   // 报告错误是否为文件或目录不存在（ErrNotExist满足）
// func IsPathSeparator(c uint8) bool                                // 报告c是否为目录分隔符
// func IsPermission(err error) bool                                 // 报告错误err是否是权限被拒绝
// func IsTimeout(err error) bool                                    // 报告错误err是否超时发生
// func Lchown(name string, uid, gid int) error                      // 修改文件的所属用户id和组id，如果文件是符号链接则修改的符号链接自身的uid和gid，如果发生错误则为*PathError类型
// func Link(oldname, newname string) error                          // 创建一newname作为oldname的硬链接，有错误则为*LinkError类型
// func LookupEnv(key string) (string, bool)                         // 检索环境变量中key对应的值，存在返回值和true，不存在返回空和false
// func Mkdir(name string, perm FileMode) error                      // 以指定权限创建name目录，有错误则为*PathError类型
// func MkdirAll(path string, perm FileMode) error                   // 以指定权限创建path目录和其必要的父目录
// func MkdirTemp(dir, pattern string) (string, error)               // 在目录dir中创建以模式pattern加末尾随机字符串的目录，返回创建的目录和错误
// func NewSyscallError(syscall string, err error) error             // 返回一个带有给定系统调用名称和错误详细信息的新的SyscallError错误，如果err为nil则返回错误为nil
// func Pipe() (r *File, w *File, err error)                         // 创建一对*File，从r读取数据返回写入w的字节
// func ReadFile(name string) ([]byte, error)                        // 读取name文件并返回所有内容和错误，成功调用返回错误为nil而不是EOF
// func Readlink(name string) (string, error)                        // 返回name符号链接的目的地，有错误则为*PathError类型
// func Remove(name string) error                                    // 删除name文件或空目录，有错误则为*PathError类型
// func RemoveAll(path string) error                                 // 删除path目录以及其包含的所有子项
// func Rename(oldpath, newpath string) error                        // 重命名oldpath为newpath，如果newpath已存在并且不是目录，Rename将替换它，有错误则为*LinkError类型
// func SameFile(fi1, fi2 FileInfo) bool                             // 报告fi1、fi2是否描述同一个文件
// func Setenv(key, value string) error                              // 环境变量设置值
// func Symlink(oldname, newname string) error                       // 创建newname文件作为oldname的符号链接，有错误则为*LinkError类型，有错误则为*PathError类型
// func TempDir() string                                             // 返回给用户存放临时文件的默认目录
// func Truncate(name string, size int64) error                      // 改变name文件的大小，如果文件是符号链接，其改变的是链接指定目标的文件大小
// func Unsetenv(key string) error                                   // unset环境变量
// func UserCacheDir() (string, error)                               // 返回用于用户特定缓存数据的默认根目录
// func UserConfigDir() (string, error)                              // 返回用于用户特定配置数据的默认根目录
// func UserHomeDir() (string, error)                                // 返回当前用户的主目录
// func WriteFile(name string, data []byte, perm FileMode) error     // 以perm权限写入data数据到name文件中，如果文件已存在则在写之前截断它

// 5.类型
// (1) DirEntry: 从目录读取的条目（使用os.ReadDir或File的ReadDir获取）
// type DirEntry = fs.DirEntry
// func ReadDir(name string) ([]DirEntry, error) // 读取name目录，按文件名排序返回所有目录条目

// (2) File: 打开的文件描述符
// type File struct {
//    // contains filtered or unexported fields
// }
// func Create(name string) (*File, error)                              // 创建或截断文件，返回文件描述符和错误。文件已存在则截断，文件不存在则以0666权限创建。
// func CreateTemp(dir, pattern string) (*File, error)                  // 在目录dir中以模式pattern创建临时文件
// func NewFile(fd uintptr, name string) *File                          // 基于给定文件描述符和名称返回一新的File
// func Open(name string) (*File, error)                                // 打开文件name用于读取，返回*File和错误
// func OpenFile(name string, flag int, perm FileMode) (*File, error)   // 以指定标志flag和指定权限perm打开文件name
// func (f *File) Chdir() error                                         // 修改当前工作目录为f，其必须是一目录
// func (f *File) Chmod(mode FileMode) error                            // 修改文件f的权限
// func (f *File) Chown(uid, gid int) error                             // 修改文件f的所属用户数字uid和用户组gid
// func (f *File) Close() error                                         // 关闭文件f，使其无法用于I/O
// func (f *File) Fd() uintptr                                          // 返回文件f的整数unix文件描述符
// func (f *File) Name() string                                         // 返回文件f的名称
// func (f *File) Read(b []byte) (n int, err error)                     // 从文件f读取最多len(b)个字节到b中，返回读取的字节数和错误，到文件末尾时返回0,io.EOF
// func (f *File) ReadAt(b []byte, off int64) (n int, err error)        // 从偏移量off位置开始读取len(b)个字节，当n<len(b)总是返回一非nil错误
// func (f *File) ReadDir(n int) ([]DirEntry, error)                    // 读取文件f关联的目录，返回部分按目录顺序的DirEntry切片
// func (f *File) ReadFrom(r io.Reader) (n int64, err error)            // 实现了io.ReaderFrom接口
// func (f *File) Readdir(n int) ([]FileInfo, error)                    // 读取文件f关联的目录，返回部分按目录顺序的FileInfo
// func (f *File) Readdirnames(n int) (names []string, err error)       // 读取文件f关联的目录，返回部分按目录顺序的文件名切片
// func (f *File) Seek(offset int64, whence int) (ret int64, err error) // 基于起始标志whence，将下一次Read和Write的偏移量设置为offset，返回新的偏移量和错误
// func (f *File) SetDeadline(t time.Time) error                        // 设置文件f读写的期限，等价于同时调用SetReadDeadline和SetWriteDeadline
// func (f *File) SetReadDeadline(t time.Time) error                    // 设置文件f的读期限
// func (f *File) SetWriteDeadline(t time.Time) error                   // 设置文件f的写期限
// func (f *File) Stat() (FileInfo, error)                              // 返回文件f的FileInfo和错误
// func (f *File) Sync() error                                          // 将文件当前内容提交到稳定存储，意味着将文件系统的最近写入的内存数据刷新到磁盘中
// func (f *File) SyscallConn() (syscall.RawConn, error)                // 返回原始文件，实现了syscall.Conn接口
// func (f *File) Truncate(size int64) error                            // 改变文件大小，不会修改I/O偏移量
// func (f *File) Write(b []byte) (n int, err error)                    // 写入len(b)字节到文件f，返回写入的字节数和错误，当n!=len(b)返回非nil错误
// func (f *File) WriteAt(b []byte, off int64) (n int, err error)       // 从off偏移量开始写入len(b)字节到文件f中，当n!=len(b)返回非nil错误
// func (f *File) WriteString(s string) (n int, err error)              // 相当于Write，只不过出传入的是字符串

// (3) FileInfo: 用于描述文件，通过函数Stat、Lstat获取
// type FileInfo = fs.FileInfo
// func Lstat(name string) (FileInfo, error) // 返回文件name的FileInfo，如果文件是符号链接，则返回的是符号链接的FileInfo
// func Stat(name string) (FileInfo, error)  // 返回文件name的FileInfo

// (4) FileMode: 描述文件的模式和权限位
// type FileMode = fs.FileMode

// (5) LinkError: 记录链接或符号链接或重命名调用时的错误和引起错误的路径
// type LinkError struct {
// 	Op  string
// 	Old string
// 	New string
// 	Err error
// }
// func (e *LinkError) Error() string
// func (e *LinkError) Unwrap() error

// (6) PathError: 记录错误以及导致它的操作和文件路径
// type PathError = fs.PathError

// (7) ProcAttr: 保存将应用于由 StartProcess 启动的新进程的属性。
// type ProcAttr struct {
// 	Dir string
// 	Env []string
// 	Files []*File
// 	Sys *syscall.SysProcAttr
// }

// (8) Process: 保存StartProcess创建进程的信息
// type Process struct {
// 	Pid int
// 	// contains filtered or unexported fields
// }
// func FindProcess(pid int) (*Process, error)                                     // 通过pid查询正在运行进程
// func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error) // 开始启动一新进程
// func (p *Process) Kill() error                                                  // 进程立刻腿粗，不会等到Process实际退出
// func (p *Process) Release() error                                               // 释放与进程p相关的任何资源，使其将来无法使用，只有在不是Wait时Release才需要调用
// func (p *Process) Signal(sig Signal) error                                      // 发送信号sig给进程
// func (p *Process) Wait() (*ProcessState, error)                                 // 等待进程退出，返回返回描述状态的ProcessState和错误

// (9) ProcessState: 保存一进程的信息，由Wait报告
// type ProcessState struct {
//    // contains filtered or unexported fields
// }
// func (p *ProcessState) ExitCode() int             // 返回进程退出的退出代码，如果进程尚未退出或被信号终止则返回-1
// func (p *ProcessState) Exited() bool              // 报告进程是否已退出
// func (p *ProcessState) Pid() int                  // 返回退出进程的进程id
// func (p *ProcessState) String() string            // 实现Stringer接口
// func (p *ProcessState) Success() bool             // 报告进程是否成功退出，unix退出状态为0
// func (p *ProcessState) Sys() interface{}          // 返回进程独立于系统的退出信息
// func (p *ProcessState) SysUsage() interface{}     // 返回独立于系统的资源使用信息
// func (p *ProcessState) SystemTime() time.Duration // 返回退出进程和其子进程的系统CPU时间
// func (p *ProcessState) UserTime() time.Duration   // 返回退出进程和其子进程的用户CPU时间

// (10) Signal: 表示操作系统信号，底层实现依赖于系统，在Unix上它是syscall.Signal
// type Signal interface {
//    String() string
//    Signal() // to distinguish from other Stringers
// }
// var (
//    Interrupt Signal = syscall.SIGINT
//    Kill      Signal = syscall.SIGKILL
// )

// (11) SyscallError: 记录特定系统调用的错误
// type SyscallError
// func (e *SyscallError) Error() string
// func (e *SyscallError) Timeout() bool // 报告错误是否表示超时
// func (e *SyscallError) Unwrap() error
