# 环境部署
```
// 配置文件
export GOROOT=/usr/local/go 
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

- GOROOT 源码包路径
- GOPATH 开发者GO开发的工作路径
- PATH 系统环境变量

---

# 优势

- 极其简单部署方式
  - 直接编译成机器码
  - 不依赖其他库
  - 直接运行部署
- 静态类型语言
  - 编译的时候检查隐藏的大部分问题
- 语言层面的并发
  - 天生的基因支持
  - 充分利用多核
- 强大的基础库支撑
  - runtime 系统调度机制（垃圾回收，调度平均分配）
  - 高效的垃圾回收GC
  - 丰富的标准库（加解密，底层，email，应用构建，进程，线程，goroutine，文本，输入输出，debug，数据结构算法，数学，文件系统，日期与时间，压缩，测试，同步机制，数据持久存储与交换，网络通信）
- 简单易学
  - 25个关键字
  - C语言简洁基因，内嵌C语言支持
  - 面向对象特性（继承，多态，封装）
  - 跨平台语言
  
```
go build serve.go 编译成包
ldd - 查看依赖
```

```golang
package main

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("goroutine", i, "...")
}

func main() {
	for i := 0; i <= 10000; i++ {
		go goFunc(i) // 开启一个并发协程
	}
	time.Sleep(time.Second)
}

```

# 适合做什么

- 云计算基础设施
  - docker, kubernetes, etcd, consul, cloudflare CDN, 七牛存储
- 基础后端软件
  - tidb, influxdb, cockroachdb
- 微服务
  - go-kit, micro, monzo bank 的 typthon, bilibili
- 互联网基础设施
  - 以太坊, hyperledger

# 不足

- 包管理，在 GitHub 上，个人账号
- 无泛化类型
- 所有 Exception 都用 Error 处理（无 try catch)
- 对 C 的降级处理，不是无缝，没有 C 降级到 asm 那么完美（序列化问题）

# 关闭 module

`go env -w GO111MODULE=off`

# Array & Slice

```go
b := [10]int{1, 2, 3, 4, 5}
fmt.Println(b)

// 关键字 range
for index, value := range b {
  fmt.Println(index, value)
}
```

```go
var slice2 []int
slice2 = make([]int, 10) // 开辟10个位置的空间
```