package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *Server) Handler(conn net.Conn) {
	// 处理业务逻辑
	fmt.Println("链接建立成功")
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
