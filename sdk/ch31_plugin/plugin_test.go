package ch31_plugin

// plugin
// plugin包实现了Go插件的加载和符号解析。
// 插件是一个带有导出函数和变量的 Go 主包，这些函数和变量已经构建：
// go build -buildmode=plugin
// 当插件第一次打开时，所有不属于程序的包的 init 函数都会被调用。主函数没有运行。一个插件只初始化一次，不能关闭。
// 目前插件仅在 Linux、FreeBSD 和 macOS 上受支持。请报告任何问题。
