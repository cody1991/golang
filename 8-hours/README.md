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

# 继承

```go
type Human struct {
	name string
	sex  string
}

type Super struct {
	Human
	level int
}

h := Human{
  name: "张三",
  sex:  "male",
}

s := Super{
  Human{
    name: "李四",
    sex:  "male",
  },
  88,
}
```

## 多态

定义 interface 和 struct

```go
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
	name  string
}

func (c *Cat) Sleep() {
	fmt.Println(c.name, "Cat Sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	color string
	name  string
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "dog"
}

func (d *Dog) Sleep() {
	fmt.Println(d.name, "Dog Sleep")
}
```

使用

```go

func showAnimal(animal Animal) {
	animal.Sleep()

	fmt.Println(animal.GetType())
	fmt.Println(animal.GetColor())
}

func main () {

  var animal Animal

	animal = &Cat{
		name:  "mimi",
		color: "black",
	}

	fmt.Println(animal)

	animal.Sleep()

	cat := Cat{
		name:  "haha",
		color: "black",
	}

	dog := Dog{
		name:  "hahaha",
		color: "while",
	}

	showAnimal(&cat)
	showAnimal(&dog)
}
```

# 类型断言

```go
value, ok := arg.(string)

if !ok {
	fmt.Println("arg is not a string")
} else {
	fmt.Println(value, "is a string")
}
```

# 变量

- 变量
  - type
    - static type 静态类型 (int, float64, string)
    - concrete type 具体类型 (interface 指向的具体类型，系统看得见的类型)
  - value

`type + value = pair` pair 对

# reflect 包

- ValueOf
  - func ValueOf(i interface) Value {...}
  - 动态获取输入参数接口中的数据的值，如果接口为空返回 0
- TypeOf
  - func TypeOf(i interface) Type {...}
  - 动态获取输入参数接口中的类型，如果接口为空返回 nil

```go
inputType := reflect.TypeOf(input)
inputValue := reflect.ValueOf(input)

fmt.Println("reflectMethod", inputType.Name(), inputValue)

// input 类型的 field 数量： NumField()
for i := 0; i < inputType.NumField(); i++ {

	// 取得第 i 个 field
	field := inputType.Field(i)
	// 取得第 i 个 field 的值
	value := inputValue.Field(i)

	fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
}

for i := 0; i < inputType.NumMethod(); i++ {
	m := inputType.Method(i)
	fmt.Printf("%s: %v", m.Name, m.Type)
}
```

# tag

```go
type Resume struct {
	Name string `json:"name" doc:"我的名字"`
	Sex  string `json:"sex" doc:"我的性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		jsonInfo := t.Field(i).Tag.Get("json")
		docInfo := t.Field(i).Tag.Get("doc")
		fmt.Println(jsonInfo)
		fmt.Println(docInfo)
	}
}

func main() {
	var re Resume

	findTag(&re)
}
```

# json

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	fmt.Println(movie)

	// 编码过程 struct => json
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("json str %s\n", jsonStr)

	// 解码过程 json => struct
	var newMove Movie
	err = json.Unmarshal(jsonStr, &newMove)

	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return
	}

	fmt.Printf("%v\n", newMove)
}
```

# goroutine

解决高CPU调度消耗和高内存占用问题

## 解决高CPU调度消耗

线程 = 多个线程(thread，内核空间) 绑定（协程调度器（轮询） 控制） 多个协程（co-routine，用户空间）

## go

co-routine => goroutine

- 内存几KB，可大量存在
- 灵活调度，可以切换

GMP => gorountine + processor + thread

## 调度器设计策略

- 复用线程
  - work stealing 机制
  - hand off 机制
- 利用并行 (GOMAXPROCS)
  - CPU核 / 2
- 抢占
  - 协程切换成本低
- 全局 G 队列

## channel & select

```go
select {
	case <-channel1:
		// 如果 channel1 成功读到数据，进行该 case 处理语句
	case channel2 <- 1:
		// 如果成功向 channel2 写入数据，进行该 case 处理语句
	default:
}
```

# gopath 弊端

- 无版本管理概念
- 无法同步一致的第三方版本库
- 无法指定第三方版本号

# go module

## 指令

- go mod init 生成 go.mod 文件
- go mod download 下载 go.mod 指定的所有依赖
- go mod tidy 整理依赖
- go mod graph 查看依赖结构
- go mod edit 编辑 go.mod 文件
- go mod vendor 导出项目所有的依赖到 vendor
- go mod verify 检验某个模块是否被修改过
- go mod why 查看为什么需要某个模块

## 配置

- GO111MODULE 是否开启 go mod 模式
- GOPROXY 下载的地址，direct表示如果找不到会回到源 GitHub 地址下载
- GOSUMOD 校验第三方库完整性
- GOPRIVATE 私有第三方库在本地放置好，不用下载校验，可以用正则

## 替换包版本

`go mod edit -replace=xxx@1.0.0=xxx@1.0.1`

## select

更多的去了解