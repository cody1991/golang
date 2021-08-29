package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string `json:"serverIp"`
	ServerPort int    `json:"serverPort"`
	Name       string `json:"name"`
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
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

func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println("NewClient 连接服务器失败")
		return
	}

	fmt.Println("NewClient 连接服务器成功")

	// 客户端业务逻辑
	select {}
}
