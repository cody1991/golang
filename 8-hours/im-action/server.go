package main

import (
	"fmt"
	"net"
	"sync"
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
	sendMsg := "[" + user.Address + "]" + user.Name + ":" + msg

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

	user := NewUser(conn)

	// 用户加入到 OnlineMap
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播用户上线消息
	s.BoardCase(user, "已上线")

	// 当前 handler 阻塞
	select {}
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
