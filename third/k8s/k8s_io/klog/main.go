package main

import (
	"flag"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	// _ = flag.Set("v", "2") // leve
	// _ = flag.Set("logtostderr", "false") // 日志写入文件而不是stdout,stderr
	// _ = flag.Set("log_dir", "./log") // 日志写入目录, 不会自动创建目录要求目录已存在
	// _ = flag.Set("log_file", "klog.log") // 日志写入文件, 与log_dir互斥
	flag.Parse()

	// 程序退出时刷新所有pending日志
	defer klog.Flush()

	klog.V(2).Info("v2 info log msg") // flag参数v>=2时才写入日志
	klog.Info("info log msg")
	klog.Warning("warning log msg")
	klog.Error("error log msg")
}

// 1.不添加任何参数
// ➜  klog git:(master) ✗ ./klog
// I0706 17:58:49.900954   79325 main.go:21] info log msg
// W0706 17:58:49.901202   79325 main.go:22] warning log msg
// E0706 17:58:49.901209   79325 main.go:23] error log msg

// 2.设置参数v
// ➜  klog git:(master) ✗ ./klog -v=2
// I0706 17:59:26.458885   79426 main.go:20] v2 info log msg
// I0706 17:59:26.459056   79426 main.go:21] info log msg
// W0706 17:59:26.459059   79426 main.go:22] warning log msg
// E0706 17:59:26.459063   79426 main.go:23] error log msg

// 3.设置参数v,logtostderr,log_dir
// ➜  klog git:(master) ✗ ./klog -v=2 -logtostderr=false -log_dir=./log
// E0706 18:01:14.328635   79786 main.go:23] error log msg
// ➜  klog git:(master) ✗ tree log
// log
// ├── klog.ERROR -> klog.macmini.yangbo.log.ERROR.20250706-180114.79786
// ├── klog.INFO -> klog.macmini.yangbo.log.INFO.20250706-180114.79786
// ├── klog.WARNING -> klog.macmini.yangbo.log.WARNING.20250706-180114.79786
// ├── klog.macmini.yangbo.log.ERROR.20250706-180114.79786
// ├── klog.macmini.yangbo.log.INFO.20250706-180114.79786
// └── klog.macmini.yangbo.log.WARNING.20250706-180114.79786

// 4.设置参数v,logtostderr,log_file
// ➜  klog git:(master) ✗ ./klog -v=2 -logtostderr=false -log_file=klog.log
// E0706 18:01:56.519084   80202 main.go:23] error log msg
// ➜  klog git:(master) ✗ cat klog.log
// Log file created at: 2025/07/06 18:01:56
// Running on machine: macmini
// Binary: Built with gc go1.22.3 for darwin/arm64
// Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
// I0706 18:01:56.518335   80202 main.go:20] v2 info log msg
// I0706 18:01:56.519077   80202 main.go:21] info log msg
// W0706 18:01:56.519081   80202 main.go:22] warning log msg
// E0706 18:01:56.519084   80202 main.go:23] error log msg

// 1) k8s组件level对应功能
// 级别	含义
// v=0	Generally useful for this to always be visible to a cluster operator.
// v=1	A reasonable default log level if you don’t want verbosity.
// v=2	Useful steady state information about the service and important log messages that may correlate to significant changes in the system.
// This is the recommended default log level for most systems.
// v=3	Extended information about changes.
// v=4	Debug level verbosity.
// v=5	Trace level verbosity.
// v=6	Display requested resources.
// v=7	Display HTTP request headers.
// v=8	Display HTTP request contents.
// v=9	Display HTTP request contents without truncation of contents.

// 2) 详细说明
// klog.V(0) - Generally useful for this to ALWAYS be visible to an operator
// Programmer errors
// Logging extra info about a panic
// CLI argument handling
// klog.V(1) - A reasonable default log level if you don’t want verbosity.
// Information about config (listening on X, watching Y)
// Errors that repeat frequently that relate to conditions that can be corrected (pod detected as unhealthy)
// klog.V(2) - Useful steady state information about the service and important log messages that may correlate to significant changes in the system. This is the recommended default log level for most systems.
// Logging HTTP requests and their exit code
// System state changing (killing pod)
// Controller state change events (starting pods)
// Scheduler log messages
// klog.V(3) - Extended information about changes
// More info about system state changes
// klog.V(4) - Debug level verbosity
// Logging in particularly thorny parts of code where you may want to come back later and check it
// klog.V(5) - Trace level verbosity
// Context to understand the steps leading up to errors and warnings
// More information for troubleshooting reported issues
