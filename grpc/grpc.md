### protoc -I=. --go_out=plugins=grpc:. data.proto 
### 命令解释
#### -I表示输入位置
#### --go_out后接的地址表示编译好的.go文件放置地址，plugins=grpc表示会将proto文件中指定的服务编译为接口代码