package signal

// os/signal: 实现了输入信号的访问，主要用于类Unix系统。

// 函数
// func Ignore(sig ...os.Signal) // 忽略提供的信号
// func Ignored(sig os.Signal) bool // 报告是否sig信号当前被忽略
// func Notify(c chan<- os.Signal, sig ...os.Signal) // 传递信号到c
// func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc) // 当列出的信号之一到达，返回的stop函数被调用，或父上下文Done通道被关闭，单号标记完成的父上下文副本，以先发生者为准
// func Reset(sig ...os.Signal) // 撤销任何先前为提供信号sig调用Notify的效果，如果没传信号则所有信号处理被重置
// func Stop(c chan<- os.Signal) // 传递停止信号到c
