#### wasm 操作步骤
##### 1、将main.go编译成 static/main.wasm 
###### 如果启用了 GO MODULES，则需要使用 go mod init 初始化模块，或设置 GO111MODULE=auto。
```shell script
GOOS=js GOARCH=wasm go build -o static/main.wasm
```
##### 2、拷贝 wasm_exec.js (JavaScript 支持文件，加载 wasm 文件时需要) 到 static 文件夹
```shell script
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
```
##### 3 安装goexec
```shell script
go get -u github.com/shurcooL/goexec 
```
##### 4 执行命令：
```shell script
 goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'
```

##### 5 以上步骤可以写在一个Makefile文件当中 直接在当前目录下运行 make 命令即可
