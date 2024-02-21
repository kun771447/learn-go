

```shell
source ~/.bash_profile 
```

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## protoc

1.  **安装 `protoc` 编译器**： 首先，您需要安装 Protocol Buffers 编译器 `protoc`。您可以从 [Protocol Buffers release 页面]() 下载适用于您系统的预编译版本，然后按照说明进行安装。

1.  **安装 gRPC 插件**： 您还需要安装用于生成 gRPC 相关代码的插件。对于 Go，您需要安装 `protoc-gen-go` 和 `protoc-gen-go-grpc` 插件。您可以使用以下命令来安装它们：

```shell
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2.  **编译 proto 文件**： 编译您的 proto 文件，生成 Go 代码和 gRPC 相关代码。例如，假设您的 `helloworld.proto` 文件位于当前目录下，您可以使用以下命令：

```shell
protoc -I . helloworld.proto --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false
```