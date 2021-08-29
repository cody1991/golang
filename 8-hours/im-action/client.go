package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string `json:"serverIp"`
	ServerPort int    `json:"serverPort"`
	Name       string `json:"name"`
	conn       net.Conn
	flag       int `json:"flag"`
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	// 返回对象
	return client
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认8888）")
}

func (c *Client) Menu() bool {
	var flag int

	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 更新用户名")
	fmt.Println("0. 退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println("请输入合法的数字范围")
		return false
	}
}

// 处理回应消息
func (c *Client) DealResponse() {
	// 一旦有数据，拷贝到 os.Stdout 标准输出上，永久阻塞监听
	io.Copy(os.Stdout, c.conn)

	// 等价于

	// for {
	// 	buf := make([]byte, 1024)
	// 	c.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (c *Client) PublicChat() {
	// 提示用户输入消息
	var msg string
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&msg)

	for msg != "exit" {
		if len(msg) != 0 {
			sendMsg := msg + "\n"
			_, err := c.conn.Write([]byte(sendMsg))

			if err != nil {
				fmt.Println("conn.Write error: ", err)
				break
			}
		}

		msg = ""
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&msg)
	}

}

func (c *Client) UpdateName() bool {

	fmt.Println("请输入用户名：")

	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"

	_, err := c.conn.Write([]byte(sendMsg))

	if err != nil {
		fmt.Println("conn.Write error: ", err)
		return false
	}

	return true
}

func (c *Client) Run() {
	for c.flag != 0 {
		for c.Menu() != true {
			// 根据不同模式处理不同的业务
		}

		switch c.flag {
		case 1:
			// 公聊
			fmt.Printf("选择公聊模式\n\n")
			c.PublicChat()
			break
		case 2:
			// 私聊
			fmt.Printf("选择私聊模式\n\n")
			break
		case 3:
			// 改名字
			fmt.Printf("选择改名字\n\n")
			c.UpdateName()
			break
		}
	}
}

func main() {

	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("NewClient 连接服务器失败")
		return
	}

	fmt.Println("NewClient 连接服务器成功")

	// 开启一个 goroutine，处理回应的消息
	go client.DealResponse()

	// 客户端业务逻辑
	client.Run()
}
