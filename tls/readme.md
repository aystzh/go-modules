#### 生成私钥和带签名的证书
```shell script
# 生成私钥
openssl genrsa -out server.key 2048
# 生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650
# 只读权限
chmod 400 server.keyei
# 根据端口号查看进程id
lsof -i:端口号 比如查看22号端口使用情况:lsof -i:22
```
#### 注意：
##### 如果在Goland当中运行，需要注意文件路径问题，直接Run会失败，在控制台中cd到指定目录用go run 运行成功