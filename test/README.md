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