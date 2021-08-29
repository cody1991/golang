package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`

	// 在线用户列表
	OnlineMap map[string]*User `json:"onlineUsers"`
	mapLock   sync.Mutex       `json:"mapLock"` // 同步

	// 消息广播 channel
	Message chan string `json:"message"`
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息
func (s *Server) BoardCase(user *User, msg string) {
	sendMsg := "[" + user.Address + "]" + user.Name + ": " + msg

	s.Message <- sendMsg
}

// 监听 message 广播消息channel的goruntine，一旦有消息发送给在线的用户

func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		// 把 msg 发送给所有在线的用户
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) Handler(conn net.Conn) {
	// 处理业务逻辑
	fmt.Println("连接建立成功")

	user := NewUser(conn, s)

	user.Online()

	isLive := make(chan bool)

	// 处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			// 合法关闭
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err: ", err)
				return
			}

			// 提出用户消息，去掉 \n
			msg := string(buf[:n-1])

			user.DoMessage(msg)

			isLive <- true
		}
	}()

	// 当前 handler 阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，激活重置定时器
		case <-time.After(time.Second * 30):
			// 超时处理
			user.SendMessage("你已经被踢下线")
			close(user.C) // 释放用的资源
			conn.Close()  // 释放连接
			return
		}
	}
}

func (s *Server) Start() {
	// socket list
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 监听 message 的 routine
	go s.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// do handler
		go s.Handler(conn)
	}

}
