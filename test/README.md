#### test命名规则
##### Test+待测试方法名
###### 测试用的参数有且只有一个，在这里是 t *testing.T
###### 基准测试(benchmark)的参数是 *testing.B，TestMain 的参数是 *testing.M 类型

#### test常用命令
##### 该 package 下所有的测试用例都会被执行:
```shell script
go test 
```
##### go test -v，-v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率
```shell script
192:test zhanghuan$ go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      github.com/aystzh/go-modules/test       0.141s


192:test zhanghuan$ go test -cover
PASS
coverage: 50.0% of statements
ok      github.com/aystzh/go-modules/test       0.291s
```
##### go test -run TestAdd -v 指定运行其中一个用例，支持通配符 *，和部分正则表达式，例如 ^、$。
```shell script
192:test zhanghuan$ go test -run TestAdd -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      github.com/aystzh/go-modules/test       0.486s
```
##### 子测试
```shell script
192:test zhanghuan$ go test -run TestMul/neg -v
=== RUN   TestMul
=== RUN   TestMul/neg
--- PASS: TestMul (0.00s)
    --- PASS: TestMul/neg (0.00s)
PASS
ok      github.com/aystzh/go-modules/test       0.468s
```
##### 基准测试
###### 函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
###### 参数为 b *testing.B
###### 执行基准测试时，需要添加 -bench 参数。
```shell script
192:benchmark zhanghuan$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/aystzh/go-modules/test/benchmark
BenchmarkHello-12               25150080                48.2 ns/op             5 B/op          1 allocs/op
BenchmarkParallel-12             5237719               245 ns/op             272 B/op          8 allocs/op
PASS
ok      github.com/aystzh/go-modules/test/benchmark     4.382s
```
###### 基准测试各个函数意义：
```shell script
type BenchmarkResult struct {
    N         int           // 迭代次数
    T         time.Duration // 基准测试花费的时间
    Bytes     int64         // 一次迭代处理的字节数
    MemAllocs uint64        // 总的分配内存的次数
    MemBytes  uint64        // 总的分配内存的字节数
}
```
##### mockgen辅助生成测试代码 
###### 使用 mockgen 生成 db_mock.go。一般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package
```shell script
install:
go get -u github.com/golang/mock/gomock
go get -u github.com/golang/mock/mockgen
generate:
192:gomock zhanghuan$ mockgen -source=db.go -destination=db_mock.go -package=main
```