### grpc-gateway示例

protoc-gen-grpc-gateway插件安装:

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2
```


由.proto文件生成go文件命令:

```shell
protoc -I ./proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./proto/hello/hello.proto
```
