## 准备
zmq4 只是 ZeroMQ 库的包装器。它不包括库本身。所以你需要安装 ZeroMQ，包括它的开发文件。

### 1.安装ZeroMQ库
```shell
# ubuntu
$ apt install libzmq4
```

### 2.检查
```shell
# 检查ZeroMQ库是否安装
$ pkg-config --modversion libzmq
4.3.1

# 检查开启CGO
$ go env CGO_ENABLED
1
```
