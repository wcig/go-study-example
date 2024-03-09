### grpc示例

protoc-gen-go-grpc插件安装: 

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


由.proto文件生成go文件命令: 

```shell
protoc --go_out=. --go-grpc_out=. ./helloworld/helloworld.proto
```
