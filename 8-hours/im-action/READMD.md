编译打包：

`go build -o server main.go server.go`

---

连接 server

`nc 127.0.0.1 8888`

--- 

命令行解析：

```go
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认8888）")
}

func main() {
	flag.Parse()
}
```

使用

```
./client --help             
Usage of ./client:
  -ip string
        设置服务器IP地址（默认127.0.0.1） (default "127.0.0.1")
  -port int
        设置服务器端口（默认8888） (default 8888)
```