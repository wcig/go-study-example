package pprof

// net/http/pprof: 提供HTTP服务器以供采集数据分析

// 引入
// import _ "net/http/pprof"
// go func() {
//    log.Println(http.ListenAndServe("localhost:6060", nil))
// }()

// 使用
// 1.使用pprof工具查看堆配置文件
// go tool pprof http://localhost:6060/debug/pprof/heap
// 2.查看30秒CPU文件
// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
// 3.查看协程块文件
// go tool pprof http://localhost:6060/debug/pprof/block
// 4.查看互斥锁争夺的持有者
// go tool pprof http://localhost:6060/debug/pprof/mutex
// 5.收集5秒的跟踪数据
// wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=5
// go tool trace trace.out
// 6.查看所有配置文件
// http://localhost:6060/debug/pprof/
// 7.教程
// https://blog.golang.org/2011/06/profiling-go-programs.html

// 函数
// func Cmdline(w http.ResponseWriter, r *http.Request)
// func Handler(name string) http.Handler
// func Index(w http.ResponseWriter, r *http.Request)
// func Profile(w http.ResponseWriter, r *http.Request)
// func Symbol(w http.ResponseWriter, r *http.Request)
// func Trace(w http.ResponseWriter, r *http.Request)
